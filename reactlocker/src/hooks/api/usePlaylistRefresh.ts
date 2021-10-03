import axios from "axios";
import { useContext } from "react";
import { useQueryClient } from "react-query";
import { AlertContext } from "../AlertContext";

export const usePlaylistRefresh = (): ((id: number) => Promise<void>) => {
	const { pushAlert } = useContext(AlertContext);
	const queryClient = useQueryClient();

	return async (id: number) => {
		return axios
			.post(`/playlist/${id}/refresh`)
			.then(() => {
				pushAlert({
					message: "Playlist was refresh successfully.",
					type: "success",
				});
				queryClient.invalidateQueries(["playlists"]);
				return;
			})
			.catch(() => {
				pushAlert({
					message: "Failed to refresh playlist",
					type: "failure",
				});
			});
	};
};
