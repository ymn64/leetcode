-- Schema
CREATE TABLE IF NOT EXISTS Employee (
  id INT,
  salary INT
);

TRUNCATE TABLE Employee;

INSERT INTO Employee (id, salary)
  VALUES ('1', '100'),
  ('2', '200'),
  ('3', '300');

-- Select the highest salary that is less than the highest salary.
SELECT
  MAX(Salary) AS SecondHighestSalary
FROM
  Employee
WHERE
  Salary < (
    SELECT
      MAX(Salary)
    FROM
      Employee);

