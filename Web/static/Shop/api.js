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
      form.reset();
    } else {
      const err = await res.text();
      resultDisplay.textContent = "Error: " + err;
    }
  } catch (err) {
    resultDisplay.textContent = "Network error: " + err.message;
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
      form.reset();
    } else {
      const err = await res.text();
      resultDisplay.textContent = "Error: " + err;
    }
  } catch (err) {
    resultDisplay.textContent = "Network error: " + err.message;
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
  });  // Add event listeners for the Role form and load dropdowns
  function setupRoleForm() {
    document
      .getElementById("roleForm")
      ?.addEventListener("submit", (e) =>
        handleFormSubmit(e, "roleForm", "webrole")
      );
  }
  
  // Initialize role form if it exists
  setupRoleForm();
  
  // Load location dropdown if it exists
  loadDropdown(
    "geolocation",
    "locationDropdown",
    "location_id",
    "location_name"
  );
}

// Attach form handler dynamically
document.addEventListener("DOMContentLoaded", () => {
  // Set up API error handling
  setupApiErrorHandling();
  
  const form = document.querySelector("form");
  if (form) {
    form.addEventListener("submit", (e) => {
      e.preventDefault();
      submitForm(form);
    });
  }
});