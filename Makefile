
DATETIME=`date +%Y%m%d_%H%M%S`

init-data:
	touch ./sql/initdb/$(DATETIME)_$(n).sql
