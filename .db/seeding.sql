-- Clear existing data
DELETE FROM schedules;
DELETE FROM train_trips;

-- Insert sample train trips and capture their generated UUIDs using CTEs,
-- then insert schedules for those trips in a single statement.
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
INSERT INTO schedules (trip_id, "time")
SELECT id, '08:00:00'::time FROM trip1 WHERE id IS NOT NULL
UNION ALL
SELECT id, '08:10:00'::time FROM trip2 WHERE id IS NOT NULL
UNION ALL
SELECT id, '08:20:00'::time FROM trip3 WHERE id IS NOT NULL
UNION ALL
SELECT id, '09:00:00'::time FROM trip4 WHERE id IS NOT NULL
UNION ALL
SELECT id, '09:15:00'::time FROM trip5 WHERE id IS NOT NULL
UNION ALL
SELECT id, '09:25:00'::time FROM trip6 WHERE id IS NOT NULL
UNION ALL
SELECT id, '10:00:00'::time FROM trip7 WHERE id IS NOT NULL
UNION ALL
SELECT id, '10:12:00'::time FROM trip8 WHERE id IS NOT NULL
UNION ALL
SELECT id, '10:22:00'::time FROM trip9 WHERE id IS NOT NULL
ON CONFLICT (trip_id, "time") DO NOTHING;