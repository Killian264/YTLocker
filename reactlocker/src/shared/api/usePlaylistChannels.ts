import { usePlaylists } from "./usePlaylists";

export const usePlaylistChannels = (): [boolean, number[]] => {
	const [loading, playlists] = usePlaylists();

	let merged: number[] = [];

	playlists.forEach((playlist) => {
		merged = [...merged, ...playlist.channels];
	});

	let filtered = merged.filter((channel, index) => {
		return merged.indexOf(channel) === index;
	});

	return [loading, filtered];
};
