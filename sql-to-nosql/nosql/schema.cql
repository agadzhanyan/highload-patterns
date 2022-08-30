DROP KEYSPACE IF EXISTS app;

CREATE KEYSPACE app WITH replication = {
    'class': 'SimpleStrategy',
    'replication_factor': 1
};

USE app;

CREATE TABLE IF NOT EXISTS activities
(
    id        uuid,
    user_id   uuid,
    timestamp timestamp,
    data      blob,
    PRIMARY KEY ((id))
) WITH comment = 'Table with all activities';