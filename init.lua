box.cfg{
    listen = '0.0.0.0:3301',
    wal_mode = 'none',
    memtx_memory = 256 * 1024 * 1024
}
box.schema.space.create('kv', {if_not_exists = true})
box.space.kv:format({
    {name = 'key', type = 'string'},
    {name = 'value', type = '*'}
})
box.space.kv:create_index('primary', {parts = {'key'}, if_not_exists = true})

local username = os.getenv("TARANTOOL_USER_NAME") or "default_user"
local password = os.getenv("TARANTOOL_USER_PASSWORD") or "default_password"

box.schema.user.create(username, {password = password, if_not_exists = true})
box.schema.user.grant(username, 'read,write,execute', 'universe', nil, {if_not_exists = true})