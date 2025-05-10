import requests
import random
import json
from datetime import datetime, timedelta

# === Configuration ===
GET_URL = "https://postgrest.kaminjitt.com/beehive?select=beehive_id,beehivestartdate"
POST_URL = "https://api.kaminjitt.com/api/harvestlog"
BEARER_TOKEN = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYXBpX3VzZXIifQ.Nx05HavMTlQfNZB3UVXe1KjAtMf_V3v2o0UT6uIA4y8"  # Replace with actual token
HEADERS = {
    "Authorization": f"Bearer {BEARER_TOKEN}",
    "Content-Type": "application/json"
}
OUTPUT_FILE = "harvestlog.json"

# === Step 1: Fetch beehive data ===
def fetch_beehives():
    response = requests.get(GET_URL, headers=HEADERS)
    response.raise_for_status()
    try:
        return response.json()
    except Exception as e:
        print(f"Error parsing JSON: {e}")
        print(f"Response status code: {response.status_code}")
        print(f"Response content: {response.text}")
        # Return fallback data for testing
        return [{"beehive_id": 1, "beehivestartdate": "2023-01-01"}]

# === Step 2: Generate random harvest log entries ===
def random_date(start_str, end_dt):
    start_dt = datetime.fromisoformat(start_str)
    if start_dt > end_dt:
        return start_dt.strftime('%Y-%m-%d')
    delta = end_dt - start_dt
    rand_date = start_dt + timedelta(seconds=random.randint(0, int(delta.total_seconds())))
    return rand_date.strftime('%Y-%m-%d')

def generate_logs(count, filename=OUTPUT_FILE):
    beehives = fetch_beehives()
    now = datetime.now()
    logs = []

    for _ in range(count):
        hive = random.choice(beehives)
        log = {
            "beehive_id": hive["beehive_id"],
            "harvestdate": random_date(hive["beehivestartdate"], now),
            "production": round(random.uniform(0.5, 10.0), 2)
        }
        logs.append(log)

    with open(filename, "w", encoding="utf-8") as f:
        json.dump(logs, f, indent=2, ensure_ascii=False)
    print(f"✅ Saved {count} logs to {filename}")

# === Step 3: Load from JSON and POST ===
def post_logs(filename=OUTPUT_FILE):
    with open(filename, "r", encoding="utf-8") as f:
        records = json.load(f)

    # Post all logs in a single request
    try:
        res = requests.post(POST_URL, headers=HEADERS, json=records)
        if res.status_code == 201:
            print(f"✅ Successfully posted {len(records)} logs in single batch")
        else:
            print(f"❌ Error {res.status_code}: {res.text}")
    except Exception as e:
        print(f"❌ Exception during batch post: {str(e)}")

# === Run ===
generate_logs(count=1)
post_logs()
