import { Playlist } from "../../shared/types";
import { usePlaylists } from "./usePlaylists";

export const usePlaylist = (id: number): [boolean, Playlist | null] => {
	const [playlistsLoading, playlists] = usePlaylists();

	let result: [boolean, Playlist | null] = [true, null];

	if (playlistsLoading || playlists === []) {
		return result;
	}

	playlists.forEach((playlist) => {
		if (playlist.id === id) {
			result = [false, playlist];
		}
	});

	return result;
};
