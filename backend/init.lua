box.cfg {
    listen = 3301,
    memtx_memory = 256
}

box.once("bootstrap", function()
    -- Создаём пользователя guest, если его ещё нет
    box.schema.user.create('guest', { if_not_exists = true })
    -- Даём права на чтение, запись и выполнение на всю базу
    box.schema.user.grant('guest', 'read,write,execute', 'universe', { if_not_exists = true })
end)