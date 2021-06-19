import axios from "axios";
import { useContext } from "react";
import { useQueryClient } from "react-query";
import { AlertContext } from "../AlertContext";

export const useCreatePlaylist = (): ((title: string, description: string) => Promise<void>) => {
	const { pushAlert } = useContext(AlertContext);
	const queryClient = useQueryClient();

	return async (title: string, description: string) => {
		return axios
			.post(`/playlist/create`, { title, description })
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
