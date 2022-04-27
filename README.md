# goflow 
go workflow framework 

# 1 feature：
  1.a workflow micro service

  2.simple and work
  
  3.use json define a workflow process
  
  4.use mysql as database and gorm library
  
# 2 Quick start

# 2.1 clone project to local disk
  git clone this repo to local disk

# 2.2 mysql database set
  create a database in mysql.
  run source db.sql for create table
  run source init.sql for init data
  edit config.json Dsn field
  
# 2.3 run project
  go mod tidy install dependency
  go run main.go
  use a web browser visit http://localhost:8080
  example.png shows how example process run.
  
# 3 test.json
  this file is a sample process config, store in table [procdef] field [resource]
  
# 3.1 start node
  only one node in process config to start process.
  
# 3.2 task node
  multiple nodes as approve node, set approvers in there node.

# 3.3 fork node
  multiple nodes as fork, set conditions to check which node next.

# 3.4 end node
  multiple nodes as end, end the process.
  
# 4 goflow database design

# 4.1 relative files
  db.sql   : this file is create tables in mysql
  init.sql : this file is init an example apply vacation process define

# 4.2 procdef
  procdef store flow config.
  key fields:
  `name` process name',
  `resource` a json string for process define,

# 4.3 procinst

  proc_inst store process instance，when user start a process,this table will insert one row.

  key fields:
  
  `name` starter name+process name,
  `procdef_id` process define ID from table procdef,
  `create_id` creator user id,
  `create_name` creator user name,
  `current_flow` current flow id from table procflow,
  `current_name` current node name,

# 4.4 procflow
  procflow store process flow，will record process sequence flow，
  
  key fields:
  
  `procdef_id` process define ID from table procdef,
  `procinst_id` process instance ID from table procinst,
  `prev_flow_id` previous flow ID,
  `prev_id` previous node id,
  `node_id` current node id,
  `node_name` current node name,
  `node_type` current node type,
  `create_id` creator user id,

# 4.5 approve
  approve store approver list，
  
  key fields:

  `procinst_id` process instance ID,
  `procflow_id` process flow ID,
  `form_id` form ID,
  `user_id` approver id,
  `user_name` approver name,
  `node_type` approve type: [and] for and sign, [or] for or sign,
  `node_id` node id,
  `data` json data store reason and other attached data,
  `deal_time` approve time,
  `deal_result` [ok] or [ng],

# 4.6 formdata
  formdata store user input form data，
  
  key fields:

  `form_id` form ID from form table,
  `procinst_id` process instance id,
  `form_name` form name,
  `task_id` task id,
  `form_data` json data store form data,

# 4.7 form
  form store a html template for user input
  
  key fields:

  `form_name` form name,
  `form` html template string,
  `form_desc` form field description,




   
