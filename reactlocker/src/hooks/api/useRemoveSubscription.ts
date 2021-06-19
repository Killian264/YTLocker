import axios from "axios";
import { useContext } from "react";
import { useQueryClient } from "react-query";
import { AlertContext } from "../AlertContext";

export const useRemoveSubscription = (): ((playlistId: number, channelId: string) => Promise<void>) => {
	const { pushAlert } = useContext(AlertContext);
	const queryClient = useQueryClient();

	return (playlistId: number, channelId: string) => {
		return axios
			.post(`/playlist/${playlistId}/unsubscribe/${channelId}`)
			.then(() => {
				pushAlert({
					message: "Subscription was removed from the playlist.",
					type: "success",
				});
				queryClient.invalidateQueries(["playlists"]);
				return;
			})
			.catch(() => {
				pushAlert({
					message: "Failed to remove subscription.",
					type: "failure",
				});
			});
	};
};
