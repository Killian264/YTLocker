import axios from "axios";
import { useContext } from "react";
import { useQueryClient } from "react-query";
import { AlertContext } from "../AlertContext";

export const usePlaylistDelete = (): ((id: number) => Promise<void>) => {
	const { pushAlert } = useContext(AlertContext);
	const queryClient = useQueryClient();

	return async (id: number) => {
		return axios
			.delete(`/playlist/${id}/delete`)
			.then(() => {
				pushAlert({
					message: "Playlist was deleted.",
					type: "success",
				});
				queryClient.invalidateQueries(["playlists"]);
				return;
			})
			.catch(() => {
				pushAlert({
					message: "Failed to delete playlist",
					type: "failure",
				});
			});
	};
};
