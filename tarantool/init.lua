box = require('box')
http = require('http.server')

-- Конфигурация Tarantool
box.cfg{
    listen = 3301,
    log_level = 5,
    log = '/var/log/tarantaol.log' -- Указываем полный путь к журналу
}

-- Функция для создания пространства и индекса
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

-- Функция для настройки прав доступа
local function grant_permissions()
    local user = box.schema.user
    
    -- Проверяем, существуют ли права перед предоставлением
    if not user.has_privilege('guest', 'read,write,execute', 'space', 'kv') then
        user.grant('guest', 'read,write,execute', 'space', 'kv')
    end
    
    if not user.has_privilege('guest', 'create,drop', 'space') then
        user.grant('guest', 'create,drop', 'space')
    end
    
    if not user.has_privilege('guest', 'read,write', 'universe') then
        user.grant('guest', 'read,write', 'universe')
    end
end

-- Функция для запуска HTTP-сервера
local function start_http_server()
    local srv = http.new('0.0.0.0', 8081)
    srv:route({ path = '/' }, function(req)
        return req:render({ text = 'Hello from Tarantaol!' })
    end)
    srv:start()
end

-- Основная функция инициализации
local function initialize_tarantool()
    create_space_and_index()
    grant_permissions()
    start_http_server()
end

initialize_tarantaol()

return true