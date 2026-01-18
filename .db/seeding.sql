-- Clear existing data
DELETE FROM schedules;
DELETE FROM train_trips;

-- Insert sample train trips and capture their generated UUIDs using CTEs (Common Table Expressions)

-- Line 1: "South"
WITH trip1 AS (
    INSERT INTO train_trips (line, station, train_type, distance)
    VALUES ('South', 'Sonoma', 'Regular', 0)
    ON CONFLICT (line, station, train_type) DO NOTHING
    RETURNING id
),
trip2 AS (
    INSERT INTO train_trips (line, station, train_type, distance)
    VALUES ('South', 'Prelude', 'Regular', 10)
    ON CONFLICT (line, station, train_type) DO NOTHING
    RETURNING id
),
trip3 AS (
    INSERT INTO train_trips (line, station, train_type, distance)
    VALUES ('South', 'Bronco', 'Regular', 20)
    ON CONFLICT (line, station, train_type) DO NOTHING
    RETURNING id
),

-- Line 2: "Central"
trip4 AS (
    INSERT INTO train_trips (line, station, train_type, distance)
    VALUES ('Central', 'Sandero', 'Regular', 0)
    ON CONFLICT (line, station, train_type) DO NOTHING
    RETURNING id
),
trip5 AS (
    INSERT INTO train_trips (line, station, train_type, distance)
    VALUES ('Central', 'Sonoma', 'Regular', 15)
    ON CONFLICT (line, station, train_type) DO NOTHING
    RETURNING id
),
trip6 AS (
    INSERT INTO train_trips (line, station, train_type, distance)
    VALUES ('Central', 'Torino', 'Regular', 25)
    ON CONFLICT (line, station, train_type) DO NOTHING
    RETURNING id
),

-- Line 3: "North"
trip7 AS (
    INSERT INTO train_trips (line, station, train_type, distance)
    VALUES ('North', 'Futurama', 'Regular', 0)
    ON CONFLICT (line, station, train_type) DO NOTHING
    RETURNING id
),
trip8 AS (
    INSERT INTO train_trips (line, station, train_type, distance)
    VALUES ('North', 'Prelude', 'Regular', 12)
    ON CONFLICT (line, station, train_type) DO NOTHING
    RETURNING id
),
trip9 AS (
    INSERT INTO train_trips (line, station, train_type, distance)
    VALUES ('North', 'Panther', 'Regular', 22)
    ON CONFLICT (line, station, train_type) DO NOTHING
    RETURNING id
)

-- Insert schedules for the trips created above
INSERT INTO schedules (trip_id, "time")
SELECT id, '08:00:00' FROM trip1 WHERE id IS NOT NULL
UNION ALL
SELECT id, '08:10:00' FROM trip2 WHERE id IS NOT NULL
UNION ALL
SELECT id, '08:20:00' FROM trip3 WHERE id IS NOT NULL
UNION ALL
SELECT id, '09:00:00' FROM trip4 WHERE id IS NOT NULL
UNION ALL
SELECT id, '09:15:00' FROM trip5 WHERE id IS NOT NULL
UNION ALL
SELECT id, '09:25:00' FROM trip6 WHERE id IS NOT NULL
UNION ALL
SELECT id, '10:00:00' FROM trip7 WHERE id IS NOT NULL
UNION ALL
SELECT id, '10:12:00' FROM trip8 WHERE id IS NOT NULL
UNION ALL
SELECT id, '10:22:00' FROM trip9 WHERE id IS NOT NULL
ON CONFLICT (trip_id, "time") DO NOTHING;