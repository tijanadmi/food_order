-- Creating the Meals table
CREATE TABLE Meals (
    MealID SERIAL PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Description TEXT,
    Price NUMERIC(10, 2) NOT NULL,
    Category VARCHAR(100),
    created_at timestamptz NOT NULL DEFAULT (now()),
	updated_at timestamptz NOT NULL DEFAULT (now())
);

-- Creating the Customers table
CREATE TABLE Customers (
    CustomerID SERIAL PRIMARY KEY,
    Email VARCHAR(255) UNIQUE NOT NULL,
    Name VARCHAR(200) NOT NULL,
    Street VARCHAR(200) NOT NULL,
    PostalCode VARCHAR(100) NOT NULL,
    City VARCHAR(100) NOT NULL,
    PhoneNumber VARCHAR(15),
    created_at timestamptz NOT NULL DEFAULT (now())
);

-- Creating the Orders table
CREATE TABLE Orders (
    OrderID SERIAL PRIMARY KEY,
    CustomerID INT NOT NULL,
    OrderDate TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    TotalAmount NUMERIC(10, 2) NOT NULL,
    created_at timestamptz NOT NULL DEFAULT (now()),
    FOREIGN KEY (CustomerID) REFERENCES Customers(CustomerID)
);

-- Creating the OrderDetails table
CREATE TABLE OrderDetails (
    OrderDetailID SERIAL PRIMARY KEY,
    OrderID INT NOT NULL,
    MealID INT NOT NULL,
    Quantity INT NOT NULL CHECK (Quantity > 0),
    Price NUMERIC(10, 2) NOT NULL,
    created_at timestamptz NOT NULL DEFAULT (now()),
    FOREIGN KEY (OrderID) REFERENCES Orders(OrderID),
    FOREIGN KEY (MealID) REFERENCES Meals(MealID)
);

-- Indexes for the Customers table
CREATE INDEX idx_customer_email ON Customers(Email);
CREATE INDEX idx_customer_city ON Customers(City);

-- Indexes for the Orders table
CREATE INDEX idx_order_customer_id ON Orders(CustomerID);
CREATE INDEX idx_order_date ON Orders(OrderDate);

-- Indexes for the OrderDetails table
CREATE INDEX idx_order_detail_order_id ON OrderDetails(OrderID);
CREATE INDEX idx_order_detail_meal_id ON OrderDetails(MealID);
CREATE INDEX idx_order_detail_order_meal ON OrderDetails(OrderID, MealID);