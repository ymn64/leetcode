-- Schema
CREATE TABLE IF NOT EXISTS Employee (
  Id INT,
  Salary INT
);

TRUNCATE TABLE Employee;

INSERT INTO Employee (id, salary)
  VALUES ('1', '100'),
  ('2', '200'),
  ('3', '300');

CREATE FUNCTION getNthHighestSalary (N INT)
  RETURNS INT READS SQL DATA
BEGIN
DECLARE
  result INT;
  SET N = N - 1;
  SELECT DISTINCT
    salary INTO result
  FROM
    Employee
  ORDER BY
    salary DESC
  LIMIT 1 OFFSET N;
  RETURN result;
END;
