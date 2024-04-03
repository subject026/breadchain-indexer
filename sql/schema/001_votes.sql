-- +goose Up
CREATE TABLE users (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    wallet_address TEXT NOT NULL
);

CREATE TABLE test_users (
    user_id UUID NOT NULL UNIQUE REFERENCES users(id)
);

CREATE TABLE projects (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    wallet_address TEXT NOT NULL
);

CREATE TABLE votes (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id),
    project_id UUID NOT NULL REFERENCES projects(id),
    value INT NOT NULL
);

CREATE TABLE slices (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    started_at TIMESTAMP NOT NULL
    
);

CREATE TABLE slice_projects (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    project_id UUID NOT NULL REFERENCES projects(id),
    slice_id UUID NOT NULL REFERENCES slices(id),
    value INT NOT NULL
);

INSERT INTO users (id, created_at, wallet_address) VALUES (uuid_in(md5(random()::text || random()::text)::cstring), NOW(), '0x09750ad360fdb7a2ee23669c4503c974d86d8694');
INSERT INTO users (id, created_at, wallet_address) VALUES (uuid_in(md5(random()::text || random()::text)::cstring), NOW(), '0xc915eC7f4CFD1C0A8Aba090F03BfaAb588aEF9B4');
INSERT INTO users (id, created_at, wallet_address) VALUES (uuid_in(md5(random()::text || random()::text)::cstring), NOW(), '0x7F85A82a2da50540412F6E526F1D00A0690a77B8');
INSERT INTO users (id, created_at, wallet_address) VALUES (uuid_in(md5(random()::text || random()::text)::cstring), NOW(), '0xBc8b85b1515E45Fb2d74333310A1d37B879732c0');
INSERT INTO users (id, created_at, wallet_address) VALUES (uuid_in(md5(random()::text || random()::text)::cstring), NOW(), '0xBBF84F9b823c42896c9723C0BE4D5f5eDe257b52');
INSERT INTO users (id, created_at, wallet_address) VALUES (uuid_in(md5(random()::text || random()::text)::cstring), NOW(), '0xD5cE086A9d4987Adf088889A520De98299E10bb5');
INSERT INTO users (id, created_at, wallet_address) VALUES (uuid_in(md5(random()::text || random()::text)::cstring), NOW(), '0x6B5C35d525D2d94c68Ab5c5AF9729092fc8771Dd');
INSERT INTO users (id, created_at, wallet_address) VALUES (uuid_in(md5(random()::text || random()::text)::cstring), NOW(), '0x4541c7745c82DF8c10bD4A58e28161534B353064');
INSERT INTO users (id, created_at, wallet_address) VALUES (uuid_in(md5(random()::text || random()::text)::cstring), NOW(), '0x0a00Fb2e074Ffaaf6c561164C6458b5C448120FC');
INSERT INTO users (id, created_at, wallet_address) VALUES (uuid_in(md5(random()::text || random()::text)::cstring), NOW(), '0xecb6ffaC05D8b4660b99B475B359FE454c77D153');

INSERT INTO test_users ( user_id )
SELECT  id
FROM    users;

INSERT INTO projects (id, created_at, updated_at, name, wallet_address) VALUES (uuid_in(md5(random()::text || random()::text)::cstring), NOW(), NOW(), 'Crypto Commons Association', '0x09750ad360fdb7a2ee23669c4503cryptocommons');
INSERT INTO projects (id, created_at, updated_at, name, wallet_address) VALUES (uuid_in(md5(random()::text || random()::text)::cstring), NOW(), NOW(), 'LaborDAO', '0x09750ad360fdb7a2ee23669c4503labordao');
INSERT INTO projects (id, created_at, updated_at, name, wallet_address) VALUES (uuid_in(md5(random()::text || random()::text)::cstring), NOW(), NOW(), 'Symbiota', '0x09750ad360fdb7a2ee23669c4503symbiota');

-- +goose Down
DROP TABLE users CASCADE;
DROP TABLE test_users CASCADE;
DROP TABLE projects CASCADE;
DROP TABLE votes CASCADE;
DROP TABLE slices CASCADE;
DROP TABLE slice_projects CASCADE;