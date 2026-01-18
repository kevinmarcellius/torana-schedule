-- Enable UUID generation
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Drop existing tables in reverse order of dependency to avoid foreign key constraints errors
DROP TABLE IF EXISTS schedules;
DROP TABLE IF EXISTS train_trips;

-- Create the train_trips table with a new UUID primary key
CREATE TABLE train_trips (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    line VARCHAR(255) NOT NULL,
    station VARCHAR(255) NOT NULL,
    train_type VARCHAR(255) NOT NULL,
    distance INT,
    -- Ensure that the combination of line, station, and train_type is unique
    UNIQUE(line, station, train_type)
);

-- Create the schedules table with a composite primary key
-- This table links directly to a train_trip via its UUID
CREATE TABLE schedules (
    trip_id UUID NOT NULL,
    "time" TIME NOT NULL,
    -- Define the foreign key relationship
    CONSTRAINT fk_train_trips
        FOREIGN KEY(trip_id)
        REFERENCES train_trips(id)
        ON DELETE CASCADE,
    -- The primary key is the combination of the trip and the time
    PRIMARY KEY(trip_id, "time")
);
