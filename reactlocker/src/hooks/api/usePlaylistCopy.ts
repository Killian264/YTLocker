import axios from "axios";
import { useContext } from "react";
import { useQueryClient } from "react-query";
import { AlertContext } from "../AlertContext";

export const usePlaylistCopy = (): ((id: number) => Promise<void>) => {
	const { pushAlert } = useContext(AlertContext);
	const queryClient = useQueryClient();

	return async (id: number) => {
		return axios
			.post(`/playlist/${id}/copy`)
			.then(() => {
				pushAlert({
					message: "Playlist was copied.",
					type: "success",
				});
				queryClient.invalidateQueries(["playlists"]);
				return;
			})
			.catch(() => {
				pushAlert({
					message: "Failed to copy playlist",
					type: "failure",
				});
			});
	};
};
