import React, { useEffect, useState } from "react";
import Admin from "~/admin/admin";
import myEnv from '../env';


type Order = {
    readonly created_at: Date;
    readonly item_packs: Array<ItemPack>
}

type ItemPack = {
    readonly items: Number;
    readonly packs: Number;
}

export default function Index() {
    const [itemsCount, setItemsCount] = useState(0);
    const [orders, setOrders] = useState(Array<Order>);
    const createOrder = () => {
        if (!itemsCount) {
            return
        }
        const requestOptions = {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ 'items_count': itemsCount })
        };
        fetch(`${myEnv.apiURL}/orders`, requestOptions)
            .then(response => response.json())
            .then((data: Order) => {
                setOrders((prev) => [data, ...prev])
            });

        setItemsCount(1)
    };

    useEffect(() => {
        fetch(`${myEnv.apiURL}/orders`)
            .then(response => response.json())
            .then((data: Array<Order>) => {
                setOrders(data)
            });
    }, [])

    return (
        <div>
            <Admin/>
            <main className="flex justify-center pt-16 pb-4">
                <form action={createOrder} >
                    <input type="number" placeholder="Add items" name="newTask" onChange={(e) => setItemsCount(Number(e.target.value))} />
                    <button type="submit">Create Order</button>
                </form>

            </main>
            <div className="flex justify-center">
                <ul>
                    {orders.map((order, index) => {
                        return (<li key={`li-order-${index}`} className="justify">
                            <span key={`li-order-span-${index}`} className="flex">
                                <span>Order {orders.length - index}:</span>
                                {order.item_packs.map((itemPack, ipIndex) => {
                                    return (
                                        <span className="itempack" key={`li-order-span-${index}-itemPack-${ipIndex}`}>
                                            {`${itemPack.items} items: ${itemPack.packs} packs`}
                                        </span>)
                                })}
                            </span>
                        </li>)
                    })}
                </ul>
            </div>
        </div>
    );
}
