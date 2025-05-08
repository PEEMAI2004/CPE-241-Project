import json
import random
from faker import Faker

fake = Faker('en_TH')
Faker.seed(0)

# Sets to ensure uniqueness
used_emails = set()
used_phones = set()

def generate_unique_email(first_name, last_name):
    base_email = f"{first_name.lower()}.{last_name.lower()}@example.co.th"
    email = base_email
    counter = 1
    while email in used_emails:
        email = f"{first_name.lower()}.{last_name.lower()}{counter}@example.co.th"
        counter += 1
    used_emails.add(email)
    return email

def generate_unique_phone():
    phone = fake.ssn()
    while phone in used_phones:
        phone = fake.ssn()
    used_phones.add(phone)
    return phone

def generate_random_person():
    first_name = fake.first_name()
    last_name = fake.last_name()
    fullname = f"{first_name} {last_name}"
    email = generate_unique_email(first_name, last_name)
    phone = generate_unique_phone()
    address = fake.address().replace('\n', ', ')
    
    return {
        "fullname": fullname,
        "email": email,
        "phone": phone,
        "address": address
    }

def generate_people_json(filename='people.json', count=10):
    people = [generate_random_person() for _ in range(count)]
    with open(filename, 'w', encoding='utf-8') as f:
        json.dump(people, f, indent=2, ensure_ascii=False)
    print(f"Generated {count} unique people and saved to {filename}")

# Run it
generate_people_json(count=10000)
