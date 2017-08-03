--找到气温升高的日期id(后一天比前一天温度升高)
select w.id from weather w, weather tmp where (datediff(w.date, tmp.date) = 1) and w.temperature > tmp.temperature;
