// useEventSource.js
import { useEffect, useState } from 'react';
import { useEvent } from 'react-use';

const useEventSource = (url) => {
    const [messages, setMessages] = useState([]);
    const [eventSource, setEventSource] = useState(null);

    useEffect(() => {
        if (!url) return;

        const es = new EventSource(url);
        setEventSource(es);

        return () => {
            if (es) {
                es.close();
            }
        };
    }, [url]);

    useEvent('message', (event) => {
        if (event && event.data) {
            setMessages((prevMessages) => [...prevMessages, event.data]);
        }
    }, eventSource);

    return messages;
};

export default useEventSource;
