CREATE TABLE Beekeeper (
  Beekeeper_ID varchar PRIMARY KEY,
  Name varchar,
  Age int,
  Email varchar,
  Phone varchar,
  Land_Area varchar,
  Location_ID varchar,
  User_ID varchar
);

CREATE TABLE BeeHive (
  BeeHive_ID varchar PRIMARY KEY,
  BeeHive_Number varchar,
  Beekeeper_ID varchar,
  BeeType_ID varchar,
  Plant_ID varchar,
  Location_ID varchar,
  BeeHive_Status varchar,
  BeeHiveStartDate timestamp,
  BeeHive_Description text
);

CREATE TABLE GeoLocation (
  Location_ID varchar PRIMARY KEY,
  Province_Name varchar,
  District_Name varchar,
  Latitude float,
  Longitude float,
  Location_Name varchar
);

CREATE TABLE BeeType (
  BeeType_ID varchar PRIMARY KEY,
  BeeType_Name varchar,
  BeeType_Description text
);

CREATE TABLE PlantType (
  Plant_ID varchar PRIMARY KEY,
  Plant_Name varchar,
  Anti_Oxidant boolean,
  Pollen boolean,
  Plant_Description text,
  Flowering_Season varchar,
  ClimateZone varchar
);

CREATE TABLE QueenBee (
  QueenBee_ID varchar PRIMARY KEY,
  BeeHive_ID varchar,
  StartDate timestamp,
  EndDate timestamp,
  Origin varchar
);

CREATE TABLE HarvestLog (
  Harvest_ID varchar PRIMARY KEY,
  BeeHive_ID varchar,
  HarvestDate timestamp,
  Production int,
  Unit varchar,
  Production_Note text
);

  CREATE TABLE HoneyStock (
    Stock_ID varchar PRIMARY KEY,
    BeeHive_ID varchar,
    Quantity int,
    Unit varchar,
    Harvest_ID varchar,
    Stock_Date timestamp
  );

CREATE TABLE Customer (
  Customer_ID varchar PRIMARY KEY,
  FullName varchar,
  Email varchar,
  Phone varchar,
  Address text,
  CreatedAt timestamp,
  User_ID varchar
);

CREATE TABLE OrderList (
  Order_ID varchar PRIMARY KEY,
  Customer_ID varchar,
  OrderDate timestamp,
  Status varchar,
  TotalAmount float
);

CREATE TABLE OrderItem (
  OrderItem_ID varchar PRIMARY KEY,
  Order_ID varchar,
  Stock_ID varchar,
  Quantity int,
  UnitPrice float
);

CREATE TABLE WebUser (
  User_ID varchar PRIMARY KEY,
  Email varchar,
  Password_Hash varchar,
  Role_ID int,
  CreatedAt timestamp
);

CREATE TABLE StaffRole (
  Role_ID int PRIMARY KEY,
  Role_Name varchar,
  Description text
);

ALTER TABLE Beekeeper ADD FOREIGN KEY (Location_ID) REFERENCES GeoLocation (Location_ID);

ALTER TABLE Beekeeper ADD FOREIGN KEY (User_ID) REFERENCES WebUser (User_ID);

ALTER TABLE BeeHive ADD FOREIGN KEY (Beekeeper_ID) REFERENCES Beekeeper (Beekeeper_ID);

ALTER TABLE BeeHive ADD FOREIGN KEY (BeeType_ID) REFERENCES BeeType (BeeType_ID);

ALTER TABLE BeeHive ADD FOREIGN KEY (Plant_ID) REFERENCES PlantType (Plant_ID);

ALTER TABLE BeeHive ADD FOREIGN KEY (Location_ID) REFERENCES GeoLocation (Location_ID);

ALTER TABLE QueenBee ADD FOREIGN KEY (BeeHive_ID) REFERENCES BeeHive (BeeHive_ID);

ALTER TABLE HarvestLog ADD FOREIGN KEY (BeeHive_ID) REFERENCES BeeHive (BeeHive_ID);

ALTER TABLE HoneyStock ADD FOREIGN KEY (BeeHive_ID) REFERENCES BeeHive (BeeHive_ID);

ALTER TABLE HoneyStock ADD FOREIGN KEY (Harvest_ID) REFERENCES HarvestLog (Harvest_ID);

ALTER TABLE Customer ADD FOREIGN KEY (User_ID) REFERENCES WebUser (User_ID);

ALTER TABLE OrderList ADD FOREIGN KEY (Customer_ID) REFERENCES Customer (Customer_ID);

ALTER TABLE OrderItem ADD FOREIGN KEY (Order_ID) REFERENCES OrderList (Order_ID);

ALTER TABLE OrderItem ADD FOREIGN KEY (Stock_ID) REFERENCES HoneyStock (Stock_ID);

ALTER TABLE WebUser ADD FOREIGN KEY (Role_ID) REFERENCES StaffRole (Role_ID);