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
        t1.value AS column1_value,
        t2.value AS column2_value,
        ABS(t1.value - t2.value) AS abs_difference
    FROM sorted_left t1
    JOIN sorted_right t2
    ON t1.row_num = t2.row_num
)
SELECT SUM(abs_difference) AS total_absolute_difference
FROM paired_data;

