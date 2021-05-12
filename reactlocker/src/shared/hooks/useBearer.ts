import { useLocalStorage } from "./useLocalStorage";

export const useBearer = (
	initial: string
): [string, (bearer: string) => void] => {
	const [cookie, setCookie] = useLocalStorage("bearer", {
		bearer: initial,
		date: Date.now(),
	});

	const setBearer = (bearer: string) => {
		let cookie = {
			bearer: bearer,
			date: Date.now(),
		};

		setCookie(cookie);
	};

	let bearer = initial;

	const date = new Date(cookie.date);

	if (!isExpired(date)) {
		bearer = cookie.bearer;
	}

	return [bearer, setBearer];
};

const isExpired = (date: Date) => {
	const twentyHours = 20 * 1000 * 60 * 60;
	const expires = Date.now() - twentyHours;
	return date < new Date(expires);
};
