const apiBase = "https://postgrest.kaminjitt.com"; // adjust as needed
const jwtToken =
  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYXBpX3VzZXIifQ.4TxmV2vnhZ5YTLw39wURDXQlzTHuAoaXHYhdTiqrNgY"; // Optional

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

  // if endpoint is harvestlog
  if (endpoint === "harvestlog") {
    submitAPI = "https://api.kaminjitt.com/api"
  } else {
    submitAPI = apiBase
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

// Attach form handler dynamically
document.addEventListener("DOMContentLoaded", () => {
  const form = document.querySelector("form");
  if (form) {
    form.addEventListener("submit", (e) => {
      e.preventDefault();
      submitForm(form);
    });
  }
// Add event listeners for the Role form and load dropdowns
function setupRoleForm() {
  document
    .getElementById("roleForm")
    ?.addEventListener("submit", (e) =>
      handleFormSubmit(e, "roleForm", "webrole")
    );
}
  loadDropdown(
    "geolocation",
    "locationDropdown",
    "location_id",
    "location_name"
  ); // load geolocation dropdown with specific columns
});

// Add event listeners for the BeeHive form and load dropdowns
function setupBeeHiveForm() {
  document
    .getElementById("beeHiveForm")
    ?.addEventListener("submit", (e) =>
      handleFormSubmit(e, "beeHiveForm", "beehive")
    );

  // Load dropdowns for relevant fields with specific columns
  loadDropdown("planttype", "plantDropdown", "plant_id", "plant_name");
  loadDropdown("beetype", "beetypeDropdown", "beetype_id", "beetype_name");
  loadDropdown(
    "beekeeper_with_name",
    "beekeeperDropdown",
    "beekeeper_id",
    "name"
  );
}

setupBeeHiveForm();

// Add event listeners for the QueenBee form and load dropdowns
function setupQueenBeeForm() {
  document
    .getElementById("queenBeeForm")
    ?.addEventListener("submit", (e) =>
      handleFormSubmit(e, "queenBeeForm", "queenbee")
    );

  // Load dropdowns for relevant fields with specific columns
  loadDropdown("beehive", "beehiveDropdown", "beehive_id", "beehive_number");
}

setupQueenBeeForm();

// Add event listeners for the orders form and load dropdowns
function setupOrdersForm() {
  document
    .getElementById("ordersForm")
    ?.addEventListener("submit", (e) =>
      handleFormSubmit(e, "ordersForm", "orders")
    );

  // Load dropdowns for relevant fields with specific columns
loadDropdown("customer", "customerDropdown", "customer_id", "fullname");
}

setupOrdersForm();