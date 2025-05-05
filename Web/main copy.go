package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

// User represents a user in our system
type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Role      int       `json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
}

// GoogleOAuthResponse represents tokens received from Google OAuth
type GoogleOAuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	IDToken      string `json:"id_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
}

// GoogleUserInfo represents user info from Google
type GoogleUserInfo struct {
	Sub           string `json:"sub"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Picture       string `json:"picture"`
}

// TokenResponse is what we send to the client
type TokenResponse struct {
	Token string `json:"token"`
}

var db *sql.DB

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to database
	dbURL := os.Getenv("DATABASE_URL")
	db, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Verify database connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to database")

	// Create router
	r := mux.NewRouter()

	// Static files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	
	// Routes
	r.HandleFunc("/", serveIndex)
	r.HandleFunc("/app", serveApp)
	r.HandleFunc("/auth/google/login", handleGoogleLogin)
	r.HandleFunc("/auth/google/callback", handleGoogleCallback)
	r.HandleFunc("/verify-token", verifyTokenHandler)

	r.HandleFunc("/api/farm/beehives", AuthMiddleware(handleFarmBeehives, 1, 2))  // Admin and Farm roles
	r.HandleFunc("/api/hr/users", AuthMiddleware(handleHRUsers, 1, 3))           // Admin and HR roles  
	r.HandleFunc("/api/shop/customers", AuthMiddleware(handleShopCustomers, 1, 4)) // Admin and Shop roles

	// Add CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, c.Handler(r)))
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}

func serveApp(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/app.html")
}

func handleFarmBeehives(w http.ResponseWriter, r *http.Request) {
    // Example implementation
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Farm beehives endpoint"))
}

func handleHRUsers(w http.ResponseWriter, r *http.Request) {
    // Example implementation
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("HR users endpoint"))
}

func handleShopCustomers(w http.ResponseWriter, r *http.Request) {
    // Example implementation
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Shop customers endpoint"))
}

func handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	// Define Google OAuth endpoint
	authURL := "https://accounts.google.com/o/oauth2/v2/auth"

	// Prepare query parameters
	params := url.Values{}
	params.Add("client_id", os.Getenv("GOOGLE_CLIENT_ID"))
	params.Add("redirect_uri", os.Getenv("REDIRECT_URI"))
	params.Add("response_type", "code")
	params.Add("scope", "email profile openid")
	params.Add("access_type", "offline")
	params.Add("prompt", "consent")

	// State parameter to prevent CSRF
	state := generateRandomState()
	params.Add("state", state)

	// Store state in cookie for validation
	cookie := http.Cookie{
		Name:     "oauth_state",
		Value:    state,
		Path:     "/",
		HttpOnly: true,
		Secure:   r.TLS != nil,
		MaxAge:   60 * 5, // 5 minutes
	}
	http.SetCookie(w, &cookie)

	// Redirect user to Google's OAuth page
	redirectURL := authURL + "?" + params.Encode()
	http.Redirect(w, r, redirectURL, http.StatusFound)
}

func handleGoogleCallback(w http.ResponseWriter, r *http.Request) {
    // Extract authorization code from query parameters
    code := r.URL.Query().Get("code")
    if code == "" {
        log.Println("No code found in callback")
        http.Error(w, "Code not found", http.StatusBadRequest)
        return
    }

    // Verify state to prevent CSRF
    stateCookie, err := r.Cookie("oauth_state")
    if err != nil || stateCookie.Value != r.URL.Query().Get("state") {
        log.Println("Invalid state parameter")
        http.Error(w, "Invalid state parameter", http.StatusBadRequest)
        return
    }

    // Exchange code for tokens
    log.Println("Exchanging code for tokens")
    tokenResponse, err := exchangeCodeForTokens(code)
    if err != nil {
        log.Printf("Error exchanging code for tokens: %v", err)
        http.Error(w, "Failed to exchange code for tokens", http.StatusInternalServerError)
        return
    }

    // Get user info using access token
    log.Println("Getting user info")
    userInfo, err := getUserInfo(tokenResponse.AccessToken)
    if err != nil {
        log.Printf("Error getting user info: %v", err)
        http.Error(w, "Failed to get user information", http.StatusInternalServerError)
        return
    }
    
    log.Printf("User authenticated: %s", userInfo.Email)

    // Check if user exists in database
    var user User
    err = db.QueryRow("SELECT User_ID, Email, Role_ID, CreatedAt FROM WebUser WHERE Email = $1", userInfo.Email).Scan(
        &user.ID, &user.Email, &user.Role, &user.CreatedAt)
    
    if err != nil {
        if err == sql.ErrNoRows {
            log.Println("User not found in database")
            http.Error(w, "User not found", http.StatusUnauthorized)
            return
        }
        log.Printf("Database error: %v", err)
        http.Error(w, "Database error", http.StatusInternalServerError)
        return
    }

    // Create JWT token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": user.ID,
        "email":   user.Email,
        "role_id": user.Role,
        "exp":     time.Now().Add(time.Hour * 24).Unix(),
    })

    // Sign the token
    tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
    if err != nil {
        log.Printf("Token signing error: %v", err)
        http.Error(w, "Token creation failed", http.StatusInternalServerError)
        return
    }

    log.Println("JWT token created successfully, redirecting to app")
    
    // Redirect to app with token
    http.Redirect(w, r, "/app?token="+url.QueryEscape(tokenString), http.StatusFound)
}

func exchangeCodeForTokens(code string) (*GoogleOAuthResponse, error) {
	tokenURL := "https://oauth2.googleapis.com/token"
	
	// Prepare form data
	data := url.Values{}
	data.Set("code", code)
	data.Set("client_id", os.Getenv("GOOGLE_CLIENT_ID"))
	data.Set("client_secret", os.Getenv("GOOGLE_CLIENT_SECRET"))
	data.Set("redirect_uri", os.Getenv("REDIRECT_URI"))
	data.Set("grant_type", "authorization_code")
	
	// Make the request
	client := &http.Client{}
	req, err := http.NewRequest("POST", tokenURL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	// Check response
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("bad status: %d, response: %s", resp.StatusCode, body)
	}
	
	// Parse response
	var tokenResponse GoogleOAuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		return nil, err
	}
	
	return &tokenResponse, nil
}

func getUserInfo(accessToken string) (*GoogleUserInfo, error) {
	// Create request to get user info
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v3/userinfo", nil)
	if err != nil {
		return nil, err
	}
	
	// Add authorization header
	req.Header.Add("Authorization", "Bearer "+accessToken)
	
	// Make the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	// Check response
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("bad status: %d, response: %s", resp.StatusCode, body)
	}
	
	// Parse response
	var userInfo GoogleUserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, err
	}
	
	return &userInfo, nil
}

func verifyTokenHandler(w http.ResponseWriter, r *http.Request) {
	// Get token from Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || len(authHeader) < 8 || authHeader[:7] != "Bearer " {
		http.Error(w, "Invalid or missing Authorization header", http.StatusUnauthorized)
		return
	}
	tokenString := authHeader[7:]

	// Parse and validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	// Token is valid
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		http.Error(w, "Invalid token claims", http.StatusUnauthorized)
		return
	}

	// Return user info
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(claims)
}

// generateRandomState creates a random state string
func generateRandomState() string {
	// In a real app, use a secure random generator
	// This is a simple implementation for demonstration
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

// AuthMiddleware is a middleware function to check JWT token
func AuthMiddleware(next http.HandlerFunc, allowedRoles ...int) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Get token from Authorization header
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" || len(authHeader) < 8 || authHeader[:7] != "Bearer " {
            http.Error(w, "Invalid or missing Authorization header", http.StatusUnauthorized)
            return
        }
        tokenString := authHeader[7:]

        // Parse and validate the token
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
            }
            return []byte(os.Getenv("JWT_SECRET")), nil
        })

        if err != nil || !token.Valid {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        // Check claims
        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok {
            http.Error(w, "Invalid token claims", http.StatusUnauthorized)
            return
        }

        // Check role if roles are specified
        if len(allowedRoles) > 0 {
            roleID, ok := claims["role_id"].(float64)
            if !ok {
                http.Error(w, "Invalid role in token", http.StatusUnauthorized)
                return
            }

            roleIntID := int(roleID)
            authorized := false
            for _, role := range allowedRoles {
                if roleIntID == role {
                    authorized = true
                    break
                }
            }

            if !authorized {
                http.Error(w, "Insufficient permissions", http.StatusForbidden)
                return
            }
        }

        // Call the wrapped handler
        next(w, r)
    }
}