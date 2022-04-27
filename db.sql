CREATE TABLE `procdef` (
  `id` varchar(16) NOT NULL,
  `company_id` varchar(16) DEFAULT NULL COMMENT 'company ID',
  `company_name` varchar(32) DEFAULT NULL COMMENT 'company name',
  `name` varchar(32) DEFAULT NULL COMMENT 'flow process name',
  `ver` varchar(10) DEFAULT NULL COMMENT 'version',
  `resource` longtext COMMENT 'flow define resource',
  `create_id` varchar(16) DEFAULT NULL COMMENT 'creator ID',
  `update_id` varchar(16) DEFAULT NULL COMMENT 'updater ID',
  `create_time` timestamp NULL DEFAULT NULL COMMENT 'create time',
  `update_time` timestamp NULL DEFAULT NULL COMMENT 'update time',
  `status` varchar(5) DEFAULT '0' COMMENT 'status',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `procinst` (
  `id` varchar(16) NOT NULL,
  `name` varchar(64) DEFAULT NULL COMMENT 'starter name+process name',
  `procdef_id` varchar(16) NOT NULL COMMENT 'process define ID',
  `create_id` varchar(16) DEFAULT NULL COMMENT 'creator id',
  `create_name` varchar(16) DEFAULT NULL COMMENT 'creator name',
  `create_dept_id` varchar(16) DEFAULT NULL COMMENT 'creator belong department id',
  `create_dept_name` varchar(32) DEFAULT NULL COMMENT 'creator belong department name',
  `create_email` varchar(64) DEFAULT NULL COMMENT 'creator email',
  `create_phone` varchar(16) DEFAULT NULL COMMENT 'creator phone',
  `current_flow` varchar(16) DEFAULT NULL COMMENT 'current flow id',
  `current_name` varchar(64) DEFAULT NULL COMMENT 'current node name',
  `start_id` varchar(64) DEFAULT NULL COMMENT 'first node id',
  `version` int DEFAULT '0' COMMENT 'version',
  `update_id` varchar(16) DEFAULT NULL COMMENT 'updater id',
  `create_time` timestamp NULL DEFAULT NULL COMMENT 'create time',
  `update_time` timestamp NULL DEFAULT NULL COMMENT 'update time',
  `status` varchar(5) DEFAULT '0' COMMENT 'status',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `procflow` (
  `id` varchar(16) NOT NULL,
  `procdef_id` varchar(16) NOT NULL COMMENT 'process define ID',
  `procinst_id` varchar(16) NOT NULL COMMENT 'process instance ID',
  `prev_flow_id` varchar(16) NOT NULL COMMENT 'previous flow ID',
  `prev_id` varchar(16) DEFAULT NULL COMMENT 'previous node id',
  `node_id` varchar(16) DEFAULT NULL COMMENT 'current node id',
  `node_name` varchar(64) DEFAULT NULL COMMENT 'current node name',
  `node_type` varchar(10) DEFAULT NULL COMMENT 'current node type',
  `create_id` varchar(16) DEFAULT NULL COMMENT 'creator id',
  `update_id` varchar(16) DEFAULT NULL COMMENT 'updater id',
  `create_time` timestamp NULL DEFAULT NULL COMMENT 'create time',
  `update_time` timestamp NULL DEFAULT NULL COMMENT 'update time',
  `status` varchar(5) DEFAULT '0' COMMENT 'status',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `approve` (
  `id` varchar(16) NOT NULL,
  `procinst_id` varchar(16) NOT NULL COMMENT 'process instance ID',
  `procflow_id` varchar(16) NOT NULL COMMENT 'process flow ID',
  `form_id` varchar(64) DEFAULT NULL COMMENT 'form ID',
  `user_id` varchar(16) DEFAULT NULL COMMENT 'approver id',
  `user_name` varchar(16) DEFAULT NULL COMMENT 'approver name',
  `user_dept_id` varchar(16) DEFAULT NULL COMMENT 'approver belong department id',
  `user_dept_name` varchar(32) DEFAULT NULL COMMENT 'approver belong department name',
  `user_email` varchar(64) DEFAULT NULL COMMENT 'approver email',
  `user_phone` varchar(16) DEFAULT NULL COMMENT 'approver phone',
  `node_type` varchar(10) DEFAULT NULL COMMENT 'approve type: and or',
  `node_id` varchar(16) DEFAULT NULL COMMENT 'node id',
  `data` text DEFAULT NULL COMMENT 'json data',
  `deal_time` timestamp NULL DEFAULT NULL COMMENT 'approve time',
  `deal_result` varchar(5) DEFAULT NULL COMMENT 'ok or ng',
  `create_id` varchar(16) DEFAULT NULL COMMENT 'creator id',
  `update_id` varchar(16) DEFAULT NULL COMMENT 'updater id',
  `create_time` timestamp NULL DEFAULT NULL COMMENT 'create time',
  `update_time` timestamp NULL DEFAULT NULL COMMENT 'update time',
  `status` varchar(5) DEFAULT '0' COMMENT 'status',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `formdata` (
  `id` varchar(64) NOT NULL,
  `form_id` varchar(64) DEFAULT NULL COMMENT 'form ID',
  `procinst_id` varchar(64) DEFAULT NULL COMMENT 'process instance id',
  `form_name` varchar(32) DEFAULT NULL COMMENT 'form name',
  `task_id` varchar(64) DEFAULT NULL COMMENT 'task id',
  `form_data` longtext COMMENT 'form data',
  `create_id` varchar(32) DEFAULT NULL COMMENT 'creator id',
  `update_id` varchar(32) DEFAULT NULL COMMENT 'updater id',
  `create_time` timestamp NULL DEFAULT NULL COMMENT 'create time',
  `update_time` timestamp NULL DEFAULT NULL COMMENT 'update time',
  `status` varchar(5) DEFAULT '0' COMMENT 'status',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `form` (
  `id` varchar(64) NOT NULL,
  `form_name` varchar(32) DEFAULT NULL COMMENT 'form name',
  `form` longtext COMMENT 'form',
  `form_desc` text DEFAULT NULL COMMENT 'form field description',
  `create_id` varchar(32) DEFAULT NULL COMMENT 'creator id',
  `update_id` varchar(32) DEFAULT NULL COMMENT 'updater id',
  `create_time` timestamp NULL DEFAULT NULL COMMENT 'create time',
  `update_time` timestamp NULL DEFAULT NULL COMMENT 'update time',
  `status` varchar(5) DEFAULT '0' COMMENT 'status',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;