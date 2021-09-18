import axios from "axios";
import { useContext } from "react";
import { useQueryClient } from "react-query";
import { Color } from "../../shared/types";
import { AlertContext } from "../AlertContext";

export const useUpdatePlaylist = (): ((
	id: number,
	title: string,
	description: string,
	color: Color
) => Promise<void>) => {
	const { pushAlert } = useContext(AlertContext);
	const queryClient = useQueryClient();

	return async (id: number, title: string, description: string, color: Color) => {
		return axios
			.post(`/playlist/${id}/update`, { title, description, color })
			.then(() => {
				pushAlert({
					message: "Playlist was updated successfully.",
					type: "success",
				});
				queryClient.invalidateQueries(["playlists"]);
				return;
			})
			.catch(() => {
				pushAlert({
					message: "Failed to update playlist.",
					type: "failure",
				});
			});
	};
};
