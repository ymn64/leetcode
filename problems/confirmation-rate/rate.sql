-- Schema
CREATE TABLE IF NOT EXISTS Signups (
  user_id INT,
  time_stamp datetime
);

CREATE TABLE IF NOT EXISTS Confirmations (
  user_id INT,
  time_stamp datetime,
  action ENUM ('confirmed', 'timeout')
);

TRUNCATE TABLE Signups;

INSERT INTO Signups (user_id, time_stamp)
  VALUES ('3', '2020-03-21 10:16:13'),
  ('7', '2020-01-04 13:57:59'),
  ('2', '2020-07-29 23:09:44'),
  ('6', '2020-12-09 10:39:37');

TRUNCATE TABLE Confirmations;

INSERT INTO Confirmations (user_id, time_stamp, action)
  VALUES ('3', '2021-01-06 03:30:46', 'timeout'),
  ('3', '2021-07-14 14:00:00', 'timeout'),
  ('7', '2021-06-12 11:57:29', 'confirmed'),
  ('7', '2021-06-13 12:58:28', 'confirmed'),
  ('7', '2021-06-14 13:59:27', 'confirmed'),
  ('2', '2021-01-22 00:00:00', 'confirmed'),
  ('2', '2021-02-28 23:59:59', 'timeout');

-- Left join on Signups.
SELECT
  Signups.user_id,
  action AS confirmation_rate
FROM
  `Signups`
  LEFT JOIN `Confirmations` ON Signups.user_id = Confirmations.user_id;

-- Use IF (which is a ternary operator really.)
SELECT
  Signups.user_id,
  IF (Confirmations.action = 'confirmed',
    1,
    0) AS confirmation_rate
FROM
  `Signups`
  LEFT JOIN `Confirmations` ON Signups.user_id = Confirmations.user_id;

-- But there is also CASE in case we have multiple cases.
SELECT
  Signups.user_id,
  CASE WHEN Confirmations.action = 'confirmed' THEN
    1
  ELSE
    0
  END AS confirmation_rate
FROM
  `Signups`
  LEFT JOIN `Confirmations` ON Signups.user_id = Confirmations.user_id;

-- Use AVG.
SELECT
  Signups.user_id,
  AVG( IF (Confirmations.action = 'confirmed', 1, 0)) AS confirmation_rate
FROM
  `Signups`
  LEFT JOIN `Confirmations` ON Signups.user_id = Confirmations.user_id
GROUP BY
  Signups.user_id;

-- Finally, use ROUND or TRUNCATE (but not FORMAT!) to get rid of the extra
-- digits.
SELECT
  Signups.user_id,
  ROUND(AVG( IF (Confirmations.action = 'confirmed', 1, 0)), 2) AS confirmation_rate
FROM
  `Signups`
  LEFT JOIN `Confirmations` ON Signups.user_id = Confirmations.user_id
GROUP BY
  Signups.user_id;

