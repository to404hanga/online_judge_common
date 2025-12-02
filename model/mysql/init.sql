CREATE DATABASE IF NOT EXISTS online_judge;

USE online_judge;

CREATE TABLE IF NOT EXISTS user (
    id bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
    username varchar(50) NOT NULL COMMENT '用户名(学号)',
    realname varchar(50) NOT NULL COMMENT '真实姓名',
    password varchar(255) NOT NULL COMMENT '密码',
    role tinyint NOT NULL DEFAULT 0 COMMENT '角色(0:普通用户,1:管理员)',
    status tinyint NOT NULL DEFAULT 0 COMMENT '状态(0:正常,1:禁用)',
    created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (id),
    UNIQUE KEY uk_username (username)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户表';

CREATE TABLE IF NOT EXISTS submission (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '提交 ID',
    competition_id BIGINT UNSIGNED NOT NULL COMMENT '比赛 ID',
    user_id BIGINT UNSIGNED NOT NULL COMMENT '用户 ID',
    problem_id BIGINT UNSIGNED NOT NULL COMMENT '题目 ID',
    code TEXT NOT NULL COMMENT '提交代码',
    stderr TEXT COMMENT '标准错误输出',
    language TINYINT NOT NULL COMMENT '提交语言 ( 0: C, 1: C++, 2: Python, 3: Java, 4: Go )',
    status TINYINT NOT NULL DEFAULT 0 COMMENT '提交状态 ( 0: 待判题, 1: 判题中, 2: 已判题 )',
    result TINYINT NOT NULL DEFAULT 0 COMMENT '判题结果 ( 0: 未判题, 1: Accepted, 2: Wrong Answer, 3: Compile Error, 4: Runtime Error, 5: Time Limit Exceeded, 6: Memory Limit Exceeded, 7: Output Limit Exceeded )',
    time_used INT NOT NULL DEFAULT -1 COMMENT '最大运行时间 ( 单位: 毫秒, -1 表示未判题 )',
    memory_used INT NOT NULL DEFAULT -1 COMMENT '最大运行内存 ( 单位: KB, -1 表示未判题 )',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',

    PRIMARY KEY (id),
    INDEX idx_competition_id_user_id_problem_id (competition_id, user_id, problem_id ASC),
    INDEX idx_created_at (created_at ASC)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='提交表';

CREATE TABLE IF NOT EXISTS problem (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '题目 ID',
    title VARCHAR(255) NOT NULL COMMENT '题目标题',
    description TEXT NOT NULL COMMENT '题目描述',
    status TINYINT NOT NULL DEFAULT 0 COMMENT '题目状态(0:未发布, 1:已发布, 2:已删除)',
    visible TINYINT NOT NULL DEFAULT 0 COMMENT '非比赛时是否可见(0:否, 1:是)',
    time_limit INT NOT NULL DEFAULT -1 COMMENT '时间限制(单位:毫秒)',
    memory_limit INT NOT NULL DEFAULT -1 COMMENT '内存限制(单位:MB)',
    creator_id BIGINT UNSIGNED NOT NULL COMMENT '创建者 ID',
    updater_id BIGINT UNSIGNED NOT NULL COMMENT '更新者 ID',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='题目表';

CREATE TABLE IF NOT EXISTS competition (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '比赛 ID',
    name VARCHAR(255) NOT NULL COMMENT '比赛名称',
    start_time DATETIME NOT NULL COMMENT '比赛开始时间',
    end_time DATETIME NOT NULL COMMENT '比赛结束时间',
    status TINYINT NOT NULL DEFAULT 0 COMMENT '比赛状态(0:未发布, 1:已发布, 2:已删除)',
    creator_id BIGINT UNSIGNED NOT NULL COMMENT '创建者 ID',
    updater_id BIGINT UNSIGNED NOT NULL COMMENT '更新者 ID',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    
    PRIMARY KEY (id),
    INDEX idx_start_time (start_time),
    INDEX idx_end_time (end_time)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='比赛表';

CREATE TABLE IF NOT EXISTS competition_user (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '比赛用户 ID',
    competition_id BIGINT UNSIGNED NOT NULL COMMENT '比赛 ID',
    user_id BIGINT UNSIGNED NOT NULL COMMENT '用户 ID',
    username varchar(50) NOT NULL COMMENT '用户名(学号)',
    realname varchar(50) NOT NULL COMMENT '真实姓名',
    status TINYINT NOT NULL DEFAULT 0 COMMENT '状态 ( 0: 正常, 1: 禁用 )',
    pass_count INT NOT NULL DEFAULT 0 COMMENT '通过题目数',
    total_time BIGINT NOT NULL DEFAULT 0 COMMENT '总耗时 ( 单位: 毫秒 )',
    retry_count INT NOT NULL DEFAULT 0 COMMENT '重试次数',

    PRIMARY KEY (id),
    UNIQUE INDEX uk_competition_user_id (competition_id, user_id),
    INDEX idx_ranking (competition_id ASC, pass_count DESC, total_time ASC) -- 按通过题目数和总耗时排序
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='比赛用户表';

CREATE TABLE IF NOT EXISTS competition_problem (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '比赛题目 ID',
    competition_id BIGINT UNSIGNED NOT NULL COMMENT '比赛 ID',
    problem_id BIGINT UNSIGNED NOT NULL COMMENT '题目 ID',
    problem_title  VARCHAR(255) NOT NULL COMMENT '题目标题',
    status TINYINT NOT NULL DEFAULT 1 COMMENT '比赛题目状态 ( 0: 禁用, 1: 启用 )',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',

    PRIMARY KEY (id),
    UNIQUE uk_competition_problem (competition_id, problem_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='比赛题目表';

-- 分割线

INSERT INTO user (username, realname, password, role, status)
VALUES
    ('admin1234567890123', 'to404hanga', '$2a$10$AAIzgJi/SZjPKOdJ2hliF./nzymKYUHaAlqbl0ugRy72m4cF2n1Pi', 1, 0), -- 密码 123456
    ('1234567890123', 'to404hanga', '$2a$10$AAIzgJi/SZjPKOdJ2hliF./nzymKYUHaAlqbl0ugRy72m4cF2n1Pi', 0, 0); -- 密码 123456

INSERT INTO competition (name, start_time, end_time, status, creator_id, updater_id)
VALUES 
    ('长时测试比赛', '2025-11-30 14:00:00', '2026-06-30 17:00:00', 0, 1, 1);

INSERT INTO problem (title, description, status, visible, time_limit, memory_limit, creator_id, updater_id) 
VALUES
    ('长时测试题目1', '这是一个长时测试题目', 1, 1, 200, 128, 1, 1);

INSERT INTO competition_problem (competition_id, problem_id, problem_title, status)
VALUES
    (1, 1, '长时测试题目1', 1);

INSERT INTO competition_user (competition_id, user_id, username, realname, status)
VALUES
    (1, 2, '1234567890123', 'to404hanga', 0);