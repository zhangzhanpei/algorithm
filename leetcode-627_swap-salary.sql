-- 交换 salary 表中的性别字段[m -> f, f -> m]
UPDATE salary SET sex = (
    CASE sex
        WHEN 'm' THEN 'f'
        ELSE 'm' END
);
