const apiFetch = async (url: string, options: RequestInit = {}) => {
    const token = localStorage.getItem('token');

    const headers: HeadersInit = {
        ...options.headers,
        Authorization: token ? `Bearer ${token}` : "",
    };

    const response = await fetch(url, { ...options, headers });

    if (response.status === 401) {
        try {
            const msg = await response.json();
            if (msg?.message?.includes('token is expired')) {
                localStorage.removeItem("token");
                window.location.href = "/";
            }
        } catch {
            localStorage.removeItem("token");
            window.location.href = "/";
        }
    }

    return response;
};

export default apiFetch
