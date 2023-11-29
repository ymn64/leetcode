-- Using LAG
SELECT DISTINCT
  num AS ConsecutiveNums
FROM
  (
    SELECT
      num,
      LAG(num, 1) OVER (
        ORDER BY
          id
      ) AS prev,
      LAG(num, 2) OVER (
        ORDER BY
          id
      ) AS prevprev
    FROM
      Logs
  ) AS sub
WHERE
  num = prev
  AND num = prevprev;

-- Using JOIN
SELECT DISTINCT
  l1.num AS ConsecutiveNums
FROM
  Logs l1
  JOIN Logs l2 ON l1.id = l2.id + 1
  JOIN Logs l3 ON l1.id = l3.id + 2
WHERE
  l1.num = l2.num
  AND l1.num = l3.num;
