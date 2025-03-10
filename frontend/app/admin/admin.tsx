import React, { useEffect, useState } from "react";
import myEnv from '../env';

export default function Admin() {
    const [packs, setPacks] = useState(String);
    const [username, setUsername] = useState(String);
    const [password, setPassword] = useState(String);

    const replacePacks = () => {
        if (!packs) {
            return
        }

        console.log(username, password)
        const packsNumbers = packs.split(',').map(packStr => Number(packStr))
        const basicAuth ='Basic ' + btoa(username + ":"+ password)

        const requestOptions = {
            method: 'PUT',
            headers: { 
                'Content-Type': 'application/json',
                'Authorization': basicAuth
             },
            body: JSON.stringify({ 'packs': packsNumbers })
        };
        fetch(`${myEnv.apiURL}/admin/packs`, requestOptions)
            .then(response => {
                console.log('replace packs response: ', response.status)
                setPacks('')
            });

    };

    return (
        <div className="flex justify-center">
            <form action={replacePacks} >
                <input type="text" placeholder="username" name="username" value={username} onChange={(e) => setUsername(e.target.value)} />
                <input type="password" placeholder="password" name="password" value={password} onChange={(e) => setPassword(e.target.value)} />
                <input type="text" placeholder="1,2,3" name="packs" value={packs} onChange={(e) => setPacks(e.target.value)} />
                <button type="submit">Replace Packs</button>
            </form>
        </div>
    )
}
