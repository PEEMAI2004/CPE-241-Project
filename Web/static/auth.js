// Handle token received from redirect
window.addEventListener('DOMContentLoaded', () => {
    // Check if there's a token in the URL (from OAuth redirect)
    const urlParams = new URLSearchParams(window.location.search);
    const token = urlParams.get('token');
    
    if (token) {
        console.log("Token received from OAuth redirect:", token.substring(0, 10) + "...");
        
        // Save token to local storage
        localStorage.setItem('token', token);
        
        // Remove token from URL to prevent bookmarking with token
        window.history.replaceState({}, document.title, "/");
        
        // Redirect to app page
        window.location.href = '/app';
        return; // Early return to prevent further execution
    }
    
    // Check if there's an error message
    const error = urlParams.get('error');
    if (error) {
        const errorElement = document.getElementById('error-message');
        errorElement.textContent = error;
        errorElement.style.display = 'block';
    }
    
    // Check if user is already logged in
    const storedToken = localStorage.getItem('token');
    if (storedToken) {
        console.log("Found existing token in localStorage");
        
        // Verify token still valid before redirecting
        fetch('/verify-token', {
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${storedToken}`
            }
        })
        .then(response => {
            if (response.ok) {
                console.log("Token verified successfully, redirecting to app");
                window.location.href = '/app';
            } else {
                console.log("Token invalid, removing it");
                localStorage.removeItem('token');
            }
        })
        .catch((error) => {
            console.error("Error verifying token:", error);
            localStorage.removeItem('token');
        });
    } else {
        console.log("No token found in localStorage");
    }
});