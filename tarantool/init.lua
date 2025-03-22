box = require('box')
http = require('http.server')

box.cfg{
    listen = 3301,
    log_level = 5,
    log = 'tarantool.log'
}

local function create_space_and_index()
    if not box.space.kv then
        box.schema.space.create('kv', {
            if_not_exists = true,
            format = {
                {name = 'key', type = 'string'},
                {name = 'value', type = '*'}
            }
        })
        
        box.space.kv:create_index('primary', {
            type = 'hash',
            parts = {'key'},
            if_not_exists = true
        })
    end
end

local function grant_permissions()
    local function safe_grant(user, privileges, object_type, object_name)
        local ok, err = pcall(box.schema.user.grant, user, privileges, object_type, object_name)
        if not ok and not tostring(err):find("already exists") then
            error(err)
        end
    end

    safe_grant('guest', 'read,write,execute', 'space', 'kv')
    safe_grant('guest', 'create,drop', 'space')
    safe_grant('guest', 'read,write', 'universe')
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