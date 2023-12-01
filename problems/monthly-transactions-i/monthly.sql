-- Schema
CREATE TABLE IF NOT EXISTS Transactions (
  id INT,
  country VARCHAR(4),
  state enum ('approved', 'declined'),
  amount INT,
  trans_date DATE
);

TRUNCATE TABLE Transactions;

INSERT INTO Transactions (id, country, state, amount, trans_date)
  VALUES ('121', 'US', 'approved', '1000', '2018-12-18'),
  ('122', 'US', 'declined', '2000', '2018-12-19'),
  ('123', 'US', 'approved', '2000', '2019-01-01'),
  ('124', 'DE', 'approved', '2000', '2019-01-07');

SELECT
  DATE_FORMAT (trans_date, '%Y-%m') AS month,
  country,
  COUNT(*) AS trans_count,
  SUM( IF (state = 'approved', 1, 0)) AS approved_count,
  SUM(amount) AS trans_total_amount,
  SUM( IF (state = 'approved', amount, 0)) AS approved_total_amount
FROM
  Transactions
GROUP BY
  DATE_FORMAT (trans_date, '%Y-%m'),
  country;

