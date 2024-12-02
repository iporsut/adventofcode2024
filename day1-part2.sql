CREATE TABLE your_table (
    column1 INT,
    column2 INT
);

WITH sorted_left AS (
    SELECT column1 AS value, ROW_NUMBER() OVER (ORDER BY column1 ASC) AS row_num
    FROM your_table
),
sorted_right AS (
    SELECT column2 AS value, ROW_NUMBER() OVER (ORDER BY column2 ASC) AS row_num
    FROM your_table
),
paired_data AS (
    SELECT
        t1.value AS c1,
        t2.value AS c2
    FROM sorted_left t1
    JOIN sorted_right t2
    ON t1.value= t2.value
)

SELECT SUM(c1)
FROM paired_data;

