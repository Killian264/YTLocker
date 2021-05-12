import { useState } from "react";

export const useLocalStorage = (key: string, initial: any) => {
	const [storedValue, setStoredValue] = useState(() => {
		try {
			const item = window.localStorage.getItem(key);
			return item ? JSON.parse(item) : initial;
		} catch (error) {
			return initial;
		}
	});

	const setValue = (value: string) => {
		setStoredValue(value);
		window.localStorage.setItem(key, JSON.stringify(value));
	};

	return [storedValue, setValue];
};
