import { Playlist } from "../../shared/types";
import { usePlaylistList } from "./usePlaylistList";

export const usePlaylist = (id: number): [boolean, Playlist | null] => {
	const [playlistsLoading, playlists] = usePlaylistList();

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
