-- decision_memory dump v1
-- generated 2026-03-18T20:00:24.910581+00:00
-- source: /scout --bootstrap /Users/kelsiandrews/gauntlet/apex-validator

CREATE TABLE IF NOT EXISTS decisions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    content TEXT NOT NULL,
    reasoning TEXT,
    positive_framing TEXT,
    status TEXT NOT NULL DEFAULT 'active'
        CHECK (status IN ('active', 'deprecated', 'superseded', 'violated')),
    source TEXT NOT NULL DEFAULT 'human'
        CHECK (source IN ('human', 'ai-discovered', 'ai-proposed')),
    superseded_by INTEGER REFERENCES decisions(id),
    created_at TEXT NOT NULL DEFAULT (datetime('now')),
    updated_at TEXT NOT NULL DEFAULT (datetime('now'))
);

CREATE TABLE IF NOT EXISTS decision_scopes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    decision_id INTEGER NOT NULL REFERENCES decisions(id) ON DELETE CASCADE,
    scope_type TEXT NOT NULL CHECK (scope_type IN ('file', 'pattern', 'tech')),
    scope_value TEXT NOT NULL
);

INSERT OR REPLACE INTO decisions (id, content, reasoning, status, source, superseded_by, created_at, updated_at) VALUES (1, 'Singleton Validate instance with struct/tag caching', 'Thread-safe, caches parsed tags per type', 'active', 'ai-discovered', NULL, '2026-03-18T20:00:24.910581+00:00', '2026-03-18T20:00:24.910581+00:00');
INSERT OR REPLACE INTO decisions (id, content, reasoning, status, source, superseded_by, created_at, updated_at) VALUES (2, 'Validators return bool, errors collected during traversal', 'Allows batch error collection without panics', 'active', 'ai-discovered', NULL, '2026-03-18T20:00:24.910581+00:00', '2026-03-18T20:00:24.910581+00:00');
INSERT OR REPLACE INTO decisions (id, content, reasoning, status, source, superseded_by, created_at, updated_at) VALUES (3, 'Tag parsing with alias expansion and OR operator', 'Supports complex rules like min=1,max=10|email', 'active', 'ai-discovered', NULL, '2026-03-18T20:00:24.910581+00:00', '2026-03-18T20:00:24.910581+00:00');
INSERT OR REPLACE INTO decisions (id, content, reasoning, status, source, superseded_by, created_at, updated_at) VALUES (4, 'Reflection-based struct traversal with CustomTypeFunc', 'Handles arbitrary Go types, pointers, interfaces', 'active', 'ai-discovered', NULL, '2026-03-18T20:00:24.910581+00:00', '2026-03-18T20:00:24.910581+00:00');
INSERT OR REPLACE INTO decisions (id, content, reasoning, status, source, superseded_by, created_at, updated_at) VALUES (5, 'Context support alongside non-context variants', 'Backward compat with simple use cases', 'active', 'ai-discovered', NULL, '2026-03-18T20:00:24.910581+00:00', '2026-03-18T20:00:24.910581+00:00');
INSERT OR REPLACE INTO decisions (id, content, reasoning, status, source, superseded_by, created_at, updated_at) VALUES (6, 'Thread-safe cache via sync.Pool and atomic.Value', 'Concurrent-safe singleton, reduced GC pressure', 'active', 'ai-discovered', NULL, '2026-03-18T20:00:24.910581+00:00', '2026-03-18T20:00:24.910581+00:00');
INSERT OR REPLACE INTO decisions (id, content, reasoning, status, source, superseded_by, created_at, updated_at) VALUES (7, 'Non-standard validators in separate package', 'Keeps core lightweight, explicit opt-in', 'active', 'ai-discovered', NULL, '2026-03-18T20:00:24.910581+00:00', '2026-03-18T20:00:24.910581+00:00');

INSERT OR REPLACE INTO decision_scopes (id, decision_id, scope_type, scope_value) VALUES (1, 1, 'tech', 'validator_instance.go');
INSERT OR REPLACE INTO decision_scopes (id, decision_id, scope_type, scope_value) VALUES (2, 2, 'tech', 'baked_in.go');
INSERT OR REPLACE INTO decision_scopes (id, decision_id, scope_type, scope_value) VALUES (3, 3, 'tech', 'cache.go');
INSERT OR REPLACE INTO decision_scopes (id, decision_id, scope_type, scope_value) VALUES (4, 4, 'tech', 'validator.go');
INSERT OR REPLACE INTO decision_scopes (id, decision_id, scope_type, scope_value) VALUES (5, 5, 'tech', 'validator_instance.go');
INSERT OR REPLACE INTO decision_scopes (id, decision_id, scope_type, scope_value) VALUES (6, 6, 'tech', 'cache.go');
INSERT OR REPLACE INTO decision_scopes (id, decision_id, scope_type, scope_value) VALUES (7, 7, 'file', 'non-standard/validators/');

