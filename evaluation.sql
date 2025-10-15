CREATE TABLE IF NOT EXISTS evaluation (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '判题 ID',
    submission_id BIGINT UNSIGNED NOT NULL COMMENT '提交 ID',
    testcase_id BIGINT UNSIGNED NOT NULL COMMENT '测试用例 ID',
    time_used INT NOT NULL DEFAULT -1 COMMENT '运行时间 ( 单位: 毫秒, -1 表示未判题 )',
    memory_used INT NOT NULL DEFAULT -1 COMMENT '运行内存 ( 单位: KB, -1 表示未判题 )',
    accepted TINYINT NOT NULL DEFAULT 0 COMMENT '是否通过 ( 0: 未判题, 1: 已通过, 2: 未通过 )',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',

    PRIMARY KEY (id),
    UNIQUE KEY uk_evaluation_submission_testcase (submission_id, testcase_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='判题表';