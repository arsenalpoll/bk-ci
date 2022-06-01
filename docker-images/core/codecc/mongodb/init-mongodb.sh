#!/bin/bash

mongoimport --host $CODECC_HOST --port 27017 --username $CODECC_USERNAME --password $CODECC_PASSWORD  --authenticationDatabase db_defect --db db_defect --collection t_checker_detail nosql/0001_codecc_db_defect_t_checker_detail_mongo.json
mongoimport --host $CODECC_HOST --port 27017 --username $CODECC_USERNAME --password $CODECC_PASSWORD  --authenticationDatabase db_task --db db_task --collection t_base_data nosql/0001_codecc_db_task_t_base_data_mongo.json
mongoimport --host $CODECC_HOST --port 27017 --username $CODECC_USERNAME --password $CODECC_PASSWORD  --authenticationDatabase db_defect --db db_defect --collection t_checker_set nosql/0002_codecc_db_defect_t_checker_set_mongo.json
mongoimport --host $CODECC_HOST --port 27017 --username $CODECC_USERNAME --password $CODECC_PASSWORD  --authenticationDatabase db_task --db db_task --collection t_tool_meta nosql/0002_codecc_db_task_t_tool_meta_mongo.json
mongoimport --host $CODECC_HOST --port 27017 --username $CODECC_USERNAME --password $CODECC_PASSWORD  --authenticationDatabase db_defect --db db_defect --collection t_red_line_meta nosql/0003_codecc_db_defect_t_red_line_meta_mongo.json
mongoimport --host $CODECC_HOST --port 27017 --username $CODECC_USERNAME --password $CODECC_PASSWORD  --authenticationDatabase db_defect --db db_defect --collection t_checker_package nosql/0004_codecc_db_defect_t_checker_package_mongo.json
mongoimport --host $CODECC_HOST --port 27017 --username $CODECC_USERNAME --password $CODECC_PASSWORD  --authenticationDatabase db_defect --db db_defect --collection t_checker_detail nosql/0006_codecc_db_defect_t_checker_detail_mongo.json
mongoimport --host $CODECC_HOST --port 27017 --username $CODECC_USERNAME --password $CODECC_PASSWORD  --authenticationDatabase db_quartz --db db_quartz --collection t_job_instance nosql/0007_codecc_db_quartz_t_job_instance_mongo.json