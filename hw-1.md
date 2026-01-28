## Создан сервис cart в HWgo/cart:

- go.mod с модулем github.com/Vietta7/HWgo/cart.
- cmd/cart/main.go HTTP сервер на :8081, JSON-RPC /rpc, передает запросы в handler, кладет токен из Authorization в context.
- internal/usecase/usecase.go — интерфейсы LomsClient, ProductClient, CartRepository, CartUsecase и структура cartUsecase, заглушки методов AddItem, DeleteItem, List, Clear, Checkout.
- internal/handler/handler.go - метод Handle, поддерживает cart.item.add, cart.item.delete, cart.list, cart.clear, cart.checkout, дергает usecase и отдает {result, error, id}.
- internal/repository/memory.go — in-memory карта user -> (sku -> count), методы AddItem, DeleteItem, ListItems, Clear.
- internal/client/loms_client.go — заглушка клиента LOMS (всегда есть сток, orderID = 1).
- internal/client/product_client.go — заглушка ProductService (фиксированное имя и цена).

## Создан сервис loms в HWgo/loms:
- go.mod с модулем github.com/Vietta7/HWgo/loms.
- cmd/loms/main.go — HTTP сервер на :8082, JSON-RPC /rpc, передает запросы в handler.
- internal/usecase/usecase.go — модели OrderStatus, OrderItem, Order; интерфейсы LomsRepository, LomsUsecase; структура lomsUsecase с заглушками OrderCreate, OrderInfo, OrderPay, OrderCancel, StockInfo.
​- internal/handler/handler.go — метод Handle для order/create, order/info, order/pay, order/cancel, stock/info.
- internal/repository/memory.go — заглушка in-memory репозитория с счетчиком nextID и пустыми реализациями методов интерфейса.

Оба сервиса собираются (go build ./...), JSON-RPC вход есть, все методы из README описаны и связаны с usecase.
​