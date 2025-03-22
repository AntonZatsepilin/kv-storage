box = require('box')
http = require('http.server')

box.cfg{
    listen = 3301,
    log_level = 5,
    log = 'tarantool.log'
}

local function create_space_and_index()
    local space_kv = box.schema.space.create('kv', {
        if_not_exists = true,
        format = {
            {name = 'key', type = 'string'},
            {name = 'value', type = '*'}
        }
    })
    
    space_kv:create_index('primary', {
        type = 'hash',
        parts = {'key'},
        if_not_exists = true
    })
end

local function grant_permissions()
    box.schema.user.grant('guest', 'read,write,execute', 'space', 'kv')
    box.schema.user.grant('guest', 'create,drop', 'space')
    box.schema.user.grant('guest', 'read,write', 'universe')
end

local function start_http_server()
    local srv = http.new('0.0.0.0', 8081)
    srv:route({ path = '/' }, function(req)
        return req:render({ text = 'Hello from Tarantool!' })
    end)
    local ok, err = pcall(srv.start, srv)
    if not ok then
        print("HTTP server error:", err)
    end
end

local function initialize_tarantool()
    create_space_and_index()
    grant_permissions()
    start_http_server()
end

initialize_tarantool()

return true