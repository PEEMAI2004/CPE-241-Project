const apiBase = "https://app.kaminjitt.com/api/postgrest"; // adjust as needed
const jwtToken =
  ""; // Optional

function parseTypes(data) {
  const parsed = {};
  for (let [key, value] of Object.entries(data)) {
    if (value === "") continue;
    if (key.endsWith("_id") || key === "beekeeper_age") {
      parsed[key] = parseInt(value);
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

  // if endpoint is harvestlog
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
        Authorization: `Bearer ${jwtToken}`,
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

// Attach form handler dynamically
document.addEventListener("DOMContentLoaded", () => {
  const form = document.querySelector("form");
  if (form) {
    form.addEventListener("submit", (e) => {
      e.preventDefault();
      submitForm(form);
    });
  }
});

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
        Authorization: `Bearer ${jwtToken}`,
      },
    });

    const items = await res.json();
    dropdown.innerHTML = "";

    items.forEach((item) => {
      const option = document.createElement("option");
      option.value = item[idColumn]; // Dynamically use the idColumn
      option.textContent = `${item[idColumn]} - ${item[nameColumn]}`; // Dynamically use the nameColumn
      dropdown.appendChild(option);
    });
  } catch (e) {
    dropdown.innerHTML = `<option>Error loading ${tableName}</option>`;
  }
}

// Add event listeners for the Beekeeper form and load dropdowns
function setupBeeKeeperForm() {
  document
    .getElementById("beekeeperForm")
    ?.addEventListener("submit", (e) =>
      handleFormSubmit(e, "beekeeperForm", "beekeeper")
    );

  // Load dropdowns for relevant fields with specific columns
  loadDropdown("webuser", "user_idDropdown", "user_id", "name");
  loadDropdown(
    "geolocation",
    "locationDropdown",
    "location_id",
    "location_name"
  );
}

setupBeeKeeperForm();

// Add event listeners for the User form and load dropdowns
function setupUserForm() {
  document
    .getElementById("userForm")
    ?.addEventListener("submit", (e) =>
      handleFormSubmit(e, "userForm", "webuser")
    );

  // Load dropdowns for relevant fields with specific columns
  loadDropdown("webrole", "role_idDropdown", "role_id", "role_name");
}

setupUserForm();

// Add event listeners for the Role form and load dropdowns
function setupRoleForm() {
  document
    .getElementById("roleForm")
    ?.addEventListener("submit", (e) =>
      handleFormSubmit(e, "roleForm", "webrole")
    );
}

setupRoleForm();
