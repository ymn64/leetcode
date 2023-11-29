SELECT
  customer_id
FROM
  `Customer`
GROUP BY
  `customer_id`
HAVING
  COUNT(DISTINCT `product_key`) = ( -- table may contain duplicates rows
    SELECT
      COUNT(`product_key`)
    FROM
      `Product`
  );
