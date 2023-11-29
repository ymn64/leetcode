-- Left join on Signups.
SELECT
  Signups.user_id,
  action as confirmation_rate
FROM
  `Signups`
  LEFT JOIN `Confirmations` ON Signups.user_id = Confirmations.user_id;

-- Use IF (which is a ternary operator really.)
SELECT
  Signups.user_id,
  IF(Confirmations.action = 'confirmed', 1, 0) AS confirmation_rate
FROM
  `Signups`
  LEFT JOIN `Confirmations` ON Signups.user_id = Confirmations.user_id;

-- But there is also CASE in case we have multiple cases.
SELECT
  Signups.user_id,
  CASE
    WHEN Confirmations.action = 'confirmed' THEN 1
    ELSE 0
  END AS confirmation_rate
FROM
  `Signups`
  LEFT JOIN `Confirmations` ON Signups.user_id = Confirmations.user_id;

-- Use AVG.
SELECT
  Signups.user_id,
  AVG(IF(Confirmations.action = 'confirmed', 1, 0)) AS confirmation_rate
FROM
  `Signups`
  LEFT JOIN `Confirmations` ON Signups.user_id = Confirmations.user_id
GROUP BY
  Signups.user_id;

-- Finally, use ROUND or TRUNCATE (but not FORMAT!) to get rid of the extra
-- digits.
SELECT
  Signups.user_id,
  ROUND(
    AVG(IF(Confirmations.action = 'confirmed', 1, 0)),
    2
  ) AS confirmation_rate
FROM
  `Signups`
  LEFT JOIN `Confirmations` ON Signups.user_id = Confirmations.user_id
GROUP BY
  Signups.user_id;
