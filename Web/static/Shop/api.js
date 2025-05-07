const apiBase = "https://app.kaminjitt.com/api/postgrest"; // adjust as needed

// Get the JWT token from localStorage
function getAuthToken() {
  return localStorage.getItem('token') || "";
}

function parseTypes(data) {
  const parsed = {};
  for (let [key, value] of Object.entries(data)) {
    if (value === "") continue;
    if (key.endsWith("_id") || key === "beekeeper_age") {
      parsed[key] = parseInt(value);
    } else if (
      key === "latitude" ||
      key === "longitude" ||
      key === "land_area" ||
      key === "production"
    ) {
      parsed[key] = parseFloat(value);
    } else {
      parsed[key] = value;
    }
  }
  return parsed;
}

async function submitForm(form) {
  const formData = new FormData(form);
  const data = parseTypes(Object.fromEntries(formData.entries()));
  const endpoint = form.dataset.endpoint;
  const resultDisplay = document.getElementById("result");

  // Determine API base URL based on endpoint
  let submitAPI;
  if (endpoint === "harvestlog") {
    submitAPI = "https://app.kaminjitt.com/api";
  } else {
    submitAPI = apiBase;
  }

  try {
    const res = await fetch(`${submitAPI}/${endpoint}`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Accept: "application/json",
        Authorization: `Bearer ${getAuthToken()}`, // Use token from localStorage
      },
      body: JSON.stringify(data),
    });

    if (res.ok) {
      resultDisplay.textContent = "Insert successful!";
      resultDisplay.style.display = "block";
      form.reset();
    } else {
      const err = await res.text();
      resultDisplay.textContent = "Error: " + err;
      resultDisplay.style.display = "block";
    }
  } catch (err) {
    resultDisplay.textContent = "Network error: " + err.message;
    resultDisplay.style.display = "block";
  }
}

// Function to extract user info from JWT token
function parseJwt(token) {
  try {
    const base64Url = token.split('.')[1];
    const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
    const jsonPayload = decodeURIComponent(atob(base64).split('').map(function(c) {
      return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
    }).join(''));

    return JSON.parse(jsonPayload);
  } catch (e) {
    console.error("Error parsing JWT token:", e);
    return {};
  }
}

// Function to get user name from user id
function getUserName(userId) {
  // Return a default value if no userId provided
  if (!userId) return "Unknown User";
  
  // Make synchronous fetch and return static string while waiting
  fetch(`${apiBase}/webuser?user_id=eq.${userId}`, {
    headers: {
      Accept: "application/json",
      Authorization: `Bearer ${getAuthToken()}`,
    },
  })
  .then(res => {
    if (!res.ok) throw new Error(`HTTP error ${res.status}`);
    return res.json();
  })
  .then(users => {
    if (users.length > 0) {
      // Update the user display field if it exists
      const userDisplay = document.getElementById('current_user');
      if (userDisplay) {
        userDisplay.value = `${users[0].user_id} - ${users[0].name || "Unknown User"}`;
      }
    }
  })
  .catch(e => {
    console.error("Error fetching user name:", e);
  });
  
  return "Loading...";  // Return immediate string while fetch happens in background
}

// Load user info from JWT token
function loadUserFromToken() {
  const token = getAuthToken();
  if (!token) {
    console.error("No authentication token found");
    return;
  }

  try {
    const userData = parseJwt(token);
    
    // Set user ID in hidden field
    const userIdField = document.getElementById('user_id');
    if (userIdField) {
      userIdField.value = userData.user_id || userData.id || '';
    }
    
    // Set user display name in disabled field
    const userDisplay = document.getElementById('current_user');
    if (userDisplay) {
      userDisplay.value = getUserName(userData.user_id || '');
    }
  } catch (e) {
    console.error("Error loading user data from token:", e);
  }
}

// Generic function to load dropdowns from any table
async function loadDropdown(tableName, dropdownId, idColumn, nameColumn) {
  const dropdown = document.getElementById(dropdownId);
  if (!dropdown) return;

  try {
    // Construct the API URL dynamically based on the table
    let url = `${apiBase}/${tableName}?select=${idColumn},${nameColumn}`;

    const res = await fetch(url, {
      headers: {
        Accept: "application/json",
        Authorization: `Bearer ${getAuthToken()}`, // Use token from localStorage
      },
    });

    if (!res.ok) {
      throw new Error(`HTTP error ${res.status}`);
    }

    const items = await res.json();
    dropdown.innerHTML = "";

    // Add empty default option
    const defaultOption = document.createElement("option");
    defaultOption.value = "";
    defaultOption.textContent = `-- Select ${tableName} --`;
    dropdown.appendChild(defaultOption);

    items.forEach((item) => {
      const option = document.createElement("option");
      option.value = item[idColumn]; // Dynamically use the idColumn
      option.textContent = `${item[idColumn]} - ${item[nameColumn]}`; // Dynamically use the nameColumn
      dropdown.appendChild(option);
    });
  } catch (e) {
    console.error(`Error loading ${tableName}:`, e);
    dropdown.innerHTML = `<option value="">Error loading ${tableName}</option>`;
  }
}

// Function to load honey stock dropdown
async function loadHoneyStockDropdown(dropdown) {
  if (!dropdown) return;

  try {
    // Construct the API URL for honeystock table
    let url = `${apiBase}/honeystock?select=stock_id,is_sold,quantity&is_sold=eq.false`;

    const res = await fetch(url, {
      headers: {
        Accept: "application/json",
        Authorization: `Bearer ${getAuthToken()}`, // Use token from localStorage
      },
    });

    if (!res.ok) {
      throw new Error(`HTTP error ${res.status}`);
    }

    const items = await res.json();
    
    // Keep the first option (if any) and clear the rest
    const firstOption = dropdown.querySelector('option:first-child');
    dropdown.innerHTML = "";
    if (firstOption) {
      dropdown.appendChild(firstOption);
    }

    items.forEach((item) => {
      const option = document.createElement("option");
      option.value = item.stock_id;
      option.textContent = `ID: ${item.stock_id} - ${item.is_sold ? 'Sold' : 'Available'} (${item.quantity % 1 === 0 ? item.quantity : item.quantity.toFixed(4)} kg)`;
      dropdown.appendChild(option);
    });
    
    // Return the promise that resolves when options are loaded
    return Promise.resolve();
  } catch (e) {
    console.error("Error loading honey stock:", e);
    dropdown.innerHTML = `<option value="">Error loading honey stock</option>`;
    return Promise.reject(e);
  }
}

// Generic form submission handler
async function handleFormSubmit(e, formId, endpoint) {
  e.preventDefault();
  const form = document.getElementById(formId);
  if (!form) return;
  
  const formData = new FormData(form);
  const data = parseTypes(Object.fromEntries(formData.entries()));
  const resultDisplay = document.getElementById("result");
  
  // Determine API base URL based on endpoint
  let submitAPI;
  if (endpoint === "harvestlog") {
    submitAPI = "https://app.kaminjitt.com/api";
  } else {
    submitAPI = apiBase;
  }

  try {
    const res = await fetch(`${submitAPI}/${endpoint}`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Accept: "application/json",
        Authorization: `Bearer ${getAuthToken()}`, // Use token from localStorage
      },
      body: JSON.stringify(data),
    });

    if (res.ok) {
      resultDisplay.textContent = "Insert successful!";
      resultDisplay.style.display = "block";
      form.reset();
    } else {
      const err = await res.text();
      resultDisplay.textContent = "Error: " + err;
      resultDisplay.style.display = "block";
    }
  } catch (err) {
    resultDisplay.textContent = "Network error: " + err.message;
    resultDisplay.style.display = "block";
  }
}

// Add event listeners for API request error handling
function setupApiErrorHandling() {
  window.addEventListener('unhandledrejection', event => {
    if (event.reason && event.reason.message && event.reason.message.includes('401')) {
      console.error('Authentication error. Redirecting to login page.');
      localStorage.removeItem('token');
      window.location.href = '/';
    }
  });
}

async function submitOrderForm(form) {
  const formData = new FormData(form);
  const user_id = parseInt(formData.get("user_id"));
  const customer_id = parseInt(formData.get("customer_id"));

  // Collect items from the form
  const itemRows = form.querySelectorAll(".item-row");
  const items = Array.from(itemRows).map(row => {
    const stockIdInput = row.querySelector('select[name="stock_id"]');
    const priceInput = row.querySelector('input[name="price"]');
    return {
      stock_id: parseInt(stockIdInput.value),
      price: parseFloat(priceInput.value),
    };
  }).filter(item => !isNaN(item.stock_id) && !isNaN(item.price));

  const payload = {
    user_id,
    customer_id,
    items
  };

  const resultDisplay = document.getElementById("result");

  try {
    const res = await fetch("https://app.kaminjitt.com/api/order", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Accept: "application/json",
        Authorization: `Bearer ${getAuthToken()}`,
      },
      body: JSON.stringify(payload),
    });

    if (res.ok) {
      resultDisplay.textContent = "Order submitted successfully!";
      resultDisplay.style.display = "block";
      form.reset();
      
      // Reset to just one empty item row
      const inputsContainer = document.getElementById("inputs");
      const itemRows = form.querySelectorAll(".item-row");
      
      // Remove all rows except the first one
      for (let i = 1; i < itemRows.length; i++) {
        itemRows[i].remove();
      }
      
      // Clear the first row's inputs
      if (itemRows.length > 0) {
        const firstRow = itemRows[0];
        firstRow.querySelector('select[name="stock_id"]').selectedIndex = 0;
        firstRow.querySelector('input[name="price"]').value = '';
        firstRow.querySelector('.remove-btn').disabled = true;
      }
      
      // Reset user-related fields from token
      loadUserFromToken();
    } else {
      const err = await res.text();
      resultDisplay.textContent = "Error: " + err;
      resultDisplay.style.display = "block";
    }
  } catch (err) {
    resultDisplay.textContent = "Network error: " + err.message;
    resultDisplay.style.display = "block";
  }
}

// Attach form handler dynamically
document.addEventListener("DOMContentLoaded", () => {
  // Set up API error handling
  setupApiErrorHandling();
});

// Add event listeners for the User form
if (document.getElementById("customerForm")) {
  const form = document.getElementById("customerForm");;
  if (form) {
    form.addEventListener("submit", (e) => {
      e.preventDefault();
      submitForm(form);
    });
  }
}

// Add event listeners for the Order form
if (document.getElementById("orderForm")) {
  const orderForm = document.getElementById("orderForm");
  if (orderForm) {
    orderForm.addEventListener("submit", (e) => {
      e.preventDefault();
      submitOrderForm(orderForm);
    });
  }
}