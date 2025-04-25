CREATE TABLE Beekeeper (
  Beekeeper_ID serial PRIMARY KEY,
  Name varchar,
  beekeeper_age int,
  Email varchar,
  Phone varchar,
  Land_Area varchar,
  Location_ID int,
  User_ID int
);

CREATE TABLE BeeHive (
  BeeHive_ID serial PRIMARY KEY,
  BeeHive_Number varchar,
  Beekeeper_ID int,
  BeeType_ID int,
  Plant_ID int,
  Location_ID int,
  BeeHive_Status varchar,
  BeeHiveStartDate timestamp,
  BeeHive_Description text
);

CREATE TABLE GeoLocation (
    serial PRIMARY KEY,
  Province_Name varchar,
  District_Name varchar,
  Latitude float,
  Longitude float,
  Location_Name varchar
);

CREATE TABLE BeeType (
  BeeType_ID serial PRIMARY KEY,
  BeeType_Name varchar,
  BeeType_Description text
);

CREATE TABLE PlantType (
  Plant_ID serial PRIMARY KEY,
  Plant_Name varchar,
  Anti_Oxidant varchar,
  Pollen varchar,
  Plant_Description text,
  Flowering_Season varchar,
  ClimateZone varchar
);

CREATE TABLE QueenBee (
  QueenBee_ID serial PRIMARY KEY,
  BeeHive_ID int,
  StartDate timestamp,
  EndDate timestamp,
  Origin varchar
);

CREATE TABLE HarvestLog (
  Harvest_ID serial PRIMARY KEY,
  BeeHive_ID int,
  HarvestDate timestamp,
  Production float,
  Unit varchar,
  Production_Note text
);

  CREATE TABLE HoneyStock (
    Stock_ID serial PRIMARY KEY,
    BeeHive_ID int,
    Quantity float,
    Unit varchar,
    Harvest_ID int,
    Stock_Date timestamp
    Is_Sold boolean default false,
  );

CREATE TABLE Customer (
  Customer_ID serial PRIMARY KEY,
  FullName varchar,
  Email varchar,
  Phone varchar,
  Address text,
  CreatedAt timestamp,
  User_ID int
);

CREATE TABLE OrderList (
  Order_ID serial PRIMARY KEY,
  Customer_ID int,
  OrderDate timestamp,
  Status varchar,
  TotalAmount float
);

CREATE TABLE OrderItem (
  OrderItem_ID serial PRIMARY KEY,
  Order_ID int,
  Stock_ID int,
  Quantity float,
  UnitPrice float
);

CREATE TABLE WebUser (
  User_ID serial PRIMARY KEY,
  Email varchar,
  Password_Hash varchar,
  Role_ID int,
  CreatedAt timestamp
);

CREATE TABLE StaffRole (
  Role_ID serial PRIMARY KEY,
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
GRANT SELECT, INSERT, UPDATE  ON Beekeeper TO api_user;
GRANT SELECT, INSERT, UPDATE  ON BeeHive TO api_user;
GRANT SELECT, INSERT, UPDATE ON GeoLocation TO api_user;
GRANT SELECT, INSERT, UPDATE  ON BeeType TO api_user;
GRANT SELECT, INSERT, UPDATE  ON PlantType TO api_user;
GRANT SELECT, INSERT, UPDATE  ON QueenBee TO api_user;
GRANT SELECT, INSERT, UPDATE  ON HarvestLog TO api_user;
GRANT SELECT, INSERT, UPDATE  ON HoneyStock TO api_user;
GRANT SELECT, INSERT, UPDATE  ON Customer TO api_user;
GRANT SELECT, INSERT, UPDATE  ON OrderList TO api_user;
GRANT SELECT, INSERT, UPDATE  ON OrderItem TO api_user;
GRANT SELECT, INSERT, UPDATE  ON WebUser TO api_user;
GRANT SELECT, INSERT, UPDATE  ON StaffRole TO api_user;
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO api_user;
