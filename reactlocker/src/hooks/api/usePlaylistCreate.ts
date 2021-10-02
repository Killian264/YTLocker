import axios from "axios";
import { useContext } from "react";
import { useQueryClient } from "react-query";
import { Color } from "../../shared/types";
import { AlertContext } from "../AlertContext";

export const usePlaylistCreate = (): ((title: string, description: string, color: Color, playlistId: number) => Promise<void>) => {
	const { pushAlert } = useContext(AlertContext);
	const queryClient = useQueryClient();

	return async (title: string, description: string, color: Color, playlistId: number) => {
		return axios
			.post(`/playlist/create`, { title, description, color, AccountID: playlistId })
			.then(() => {
				pushAlert({
					message: "Playlist was created.",
					type: "success",
				});
				queryClient.invalidateQueries(["playlists"]);
				return;
			})
			.catch(() => {
				pushAlert({
					message: "Failed to create playlist",
					type: "failure",
				});
			});
	};
};
