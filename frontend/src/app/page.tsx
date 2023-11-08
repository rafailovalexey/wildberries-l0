"use client"

import {useState} from "react";
import {AxiosHeaders, default as axios} from "axios";

export default function Home() {
    const [inputValue, setInputValue] = useState("");
    const [data, setData] = useState("{}");

    const handleRequest = async () => {
        try {
            const response = await axios.get(
                `http://development.emptyhopes.ru:3000/v1/orders/${inputValue}`, {
                    headers: {
                        "authentication-token": "token",
                    } as AxiosHeaders,
                }
            );

            setData(response.data);
        }
        catch (error) {
            setData(error.response.data);
        }
    }

    return (
        <main className="flex min-h-screen flex-col items-center justify-between p-24">
            <div className="z-10 max-w-5xl w-full items-center justify-between font-mono text-sm lg:flex">
                <div
                    className="fixed bottom-0 left-0 flex h-48 w-full items-end justify-center bg-gradient-to-t from-white via-white dark:from-black dark:via-black lg:static lg:h-auto lg:w-auto lg:bg-none"></div>
            </div>

            <div className="block" style={{width: "370px"}}>
                <label htmlFor="id" className="block text-sm font-medium leading-6 text-gray-900">
                    UUID
                </label>
                <div className="relative mt-2 rounded-md shadow-sm w-360">
                    <div className="pointer-events-none absolute inset-y-0 left-0 flex items-center pl-3">
                        <span className="text-gray-500 sm:text-sm">UUID</span>
                    </div>
                    <input
                        type="text"
                        name="id"
                        id="id"
                        className="w-full block rounded-md border-0 py-3 pl-14 pr-4 text-gray-900 ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-1 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                        placeholder="e24d4915-7433-4870-85ea-024e2422e026"
                        value={inputValue}
                        onChange={(e) => setInputValue(e.target.value)}
                    />
                </div>

                <button className="w-full mt-4 bg-gray-300 hover-bg-gray-400 text-gray-800 font-bold py-2 px-4 rounded"
                        onClick={handleRequest}>
                    <span>Request</span>
                </button>
            </div>

            <pre className="mt-8">{JSON.stringify(data, null, 2)}</pre>

            <div
                className="relative flex place-items-center before:absolute before:h-[300px] before:w-[480px] before:-translate-x-1/2 before:rounded-full before:bg-gradient-radial before:from-white before:to-transparent before:blur-2xl before:content-[''] after:absolute after:-z-20 after:h-[180px] after:w-[240px] after:translate-x-1/3 after:bg-gradient-conic after:from-sky-200 after:via-blue-200 after:blur-2xl after:content-[''] before:dark:bg-gradient-to-br before:dark:from-transparent before:dark:to-blue-700 before:dark:opacity-10 after:dark:from-sky-900 after:dark:via-[#0141ff] after:dark:opacity-40 before:lg:h-[360px] z-[-1]"/>
            <div className="mb-32 grid text-center lg:max-w-5xl lg:w-full lg:mb-0 lg:grid-cols-4 lg:text-left"/>
        </main>
    )
}
