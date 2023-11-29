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
      Employee
  );
