CREATE TABLE GeoLocation (
    Location_ID INT PRIMARY KEY,
    Province_Name VARCHAR(100),
    District_Name VARCHAR(100),
    Latitude DECIMAL(9,6),
    Longitude DECIMAL(9,6),
    Location_Name VARCHAR(100)
);

CREATE TABLE BeeKeeper (
    BeeKeeper_ID INT PRIMARY KEY,
    BeeKeeper_Age INT,
    BeeKeeper_Name VARCHAR(100),
    email VARCHAR(100),
    phone VARCHAR(20),
    Land_Area DECIMAL(10,2),
    Location_ID INT,
    FOREIGN KEY (Location_ID) REFERENCES GeoLocation(Location_ID)
);

CREATE TABLE BeeType (
    BeeType_ID INT PRIMARY KEY,
    BeeType_Name VARCHAR(100),
    BeeType_Description TEXT
);

CREATE TABLE PlantType (
    Plant_ID INT PRIMARY KEY,
    Plant_Name VARCHAR(100),
    Anti_Oxidant VARCHAR(100),
    Pollen VARCHAR(100),
    Plant_Description TEXT,
    Flowering_Season VARCHAR(100),
    ClimateZone VARCHAR(100)
);

CREATE TABLE Customer (
    Customer_ID INT PRIMARY KEY,
    FullName VARCHAR(100),
    Email VARCHAR(100),
    Phone VARCHAR(20),
    Address TEXT,
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE BeeHive (
    BeeHive_ID INT PRIMARY KEY,
    BeeHive_Number VARCHAR(50),
    BeeHive_Status VARCHAR(50),
    BeeHiveStartDate DATE,
    BeeHive_Description TEXT,
    Plant_ID INT,
    BeeType_ID INT,
    BeeKeeper_ID INT,
    FOREIGN KEY (Plant_ID) REFERENCES PlantType(Plant_ID),
    FOREIGN KEY (BeeType_ID) REFERENCES BeeType(BeeType_ID),
    FOREIGN KEY (BeeKeeper_ID) REFERENCES BeeKeeper(BeeKeeper_ID)
);

CREATE TABLE QueenBee (
    QueenBee_ID INT PRIMARY KEY,
    BeeHive_ID INT,
    StartDate DATE,
    EndDate DATE,
    Origin VARCHAR(100),
    FOREIGN KEY (BeeHive_ID) REFERENCES BeeHive(BeeHive_ID)
);

CREATE TABLE HarvestLog (
    Harvest_ID INT PRIMARY KEY,
    HarvestDate DATE,
    BeeHive_ID INT,
    Production DECIMAL(10,2),
    Unit VARCHAR(20),
    Production_Note TEXT,
    FOREIGN KEY (BeeHive_ID) REFERENCES BeeHive(BeeHive_ID)
);

CREATE TABLE HoneyStock (
    Stock_ID INT PRIMARY KEY,
    BeeHive_ID INT,
    Quantity DECIMAL(10,2),
    Unit VARCHAR(20),
    Harvest_ID INT,
    Stock_Date DATE,
    FOREIGN KEY (BeeHive_ID) REFERENCES BeeHive(BeeHive_ID),
    FOREIGN KEY (Harvest_ID) REFERENCES HarvestLog(Harvest_ID)
);

CREATE TABLE Orders (
    Order_ID INT PRIMARY KEY,
    Customer_ID INT,
    OrderDate TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    Status VARCHAR(50),
    TotalAmount DECIMAL(10,2),
    FOREIGN KEY (Customer_ID) REFERENCES Customer(Customer_ID)
);

CREATE TABLE OrderItem (
    OrderItem_ID INT PRIMARY KEY,
    Order_ID INT,
    Stock_ID INT,
    Quantity DECIMAL(10,2),
    UnitPrice DECIMAL(10,2),
    FOREIGN KEY (Order_ID) REFERENCES Orders(Order_ID),
    FOREIGN KEY (Stock_ID) REFERENCES HoneyStock(Stock_ID)
);
