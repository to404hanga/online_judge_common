CREATE TABLE IF NOT EXISTS problem (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '题目 ID',
    title VARCHAR(255) NOT NULL COMMENT '题目标题',
    description_url VARCHAR(255) NOT NULL COMMENT '题目 url',
    testcase_zip_url VARCHAR(255) NOT NULL COMMENT '测试用例 zip 文件 url',
    status TINYINT NOT NULL DEFAULT 0 COMMENT '题目状态(0:未发布, 1:已发布, 2:已删除)',
    visible TINYINT NOT NULL DEFAULT 0 COMMENT '是否仅比赛时可见(0:否, 1:是)',
    time_limit INT NOT NULL DEFAULT -1 COMMENT '时间限制(单位:毫秒)',
    memory_limit INT NOT NULL DEFAULT -1 COMMENT '内存限制(单位:KB)',
    creator_id BIGINT UNSIGNED NOT NULL COMMENT '创建者 ID',
    updater_id BIGINT UNSIGNED NOT NULL COMMENT '更新者 ID',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    
    PRIMARY KEY (id),
    UNIQUE INDEX uk_description_url (description_url),
    UNIQUE INDEX uk_testcase_zip_url (testcase_zip_url)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='题目表';

