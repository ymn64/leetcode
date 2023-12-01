-- Schema
CREATE TABLE IF NOT EXISTS Employee (
  id INT,
  name VARCHAR(255),
  department VARCHAR(255),
  managerId INT
);

TRUNCATE TABLE Employee;

INSERT INTO Employee (id, name, department, managerId)
  VALUES ('101', 'John', 'A', 'None'),
  ('102', 'Dan', 'A', '101'),
  ('103', 'James', 'A', '101'),
  ('104', 'Amy', 'A', '101'),
  ('105', 'Anne', 'A', '101'),
  ('106', 'Ron', 'B', '101');

SELECT
  m.name
FROM
  Employee e
  JOIN Employee m ON e.managerId = m.id
GROUP BY
  m.id,
  m.name
HAVING
  COUNT(e.id) > 4;

